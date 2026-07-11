package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *Service) MoveSymbols(
	directory string,
	packagePath string,
	symbols []string,
	filePath string,
	targetPackagePath string,
	targetFile string,
	create bool,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if packagePath == targetPackagePath {
		return failValidation(r, "source and target package are the same")
	}

	if filePath != "" && len(symbols) > 0 {
		return failValidation(r, "pass symbols or file, not both")
	}

	if filePath == "" && len(symbols) == 0 {
		return failValidation(r, "pass symbols or file")
	}

	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, e
	}

	p := findPackage(all, packagePath)

	if p == nil {
		return failValidation(
			r,
			fmt.Sprintf("package not found: %s", packagePath),
		)
	}

	if filePath != "" {
		full := filePath

		if !filepath.IsAbs(full) {
			full = filepath.Join(directory, filePath)
		}

		file := findSyntaxFile(set, p, full)

		if file == nil {
			return failValidation(
				r,
				fmt.Sprintf(
					"file not found in %s: %s",
					packagePath,
					filePath,
				),
			)
		}

		symbols, e = expandFileSymbols(file)

		if e != nil {
			return failValidation(r, e.Error())
		}
	}

	unique := make(map[string]bool)

	for _, name := range symbols {
		if unique[name] {
			return failValidation(
				r,
				fmt.Sprintf("duplicate symbol: %s", name),
			)
		}

		unique[name] = true
	}

	entries, message := buildMoveEntries(
		all,
		p,
		packagePath,
		symbols,
		targetFile,
	)

	if message != "" {
		return failValidation(r, message)
	}

	target := findPackage(all, targetPackagePath)

	if target == nil && !create {
		return failValidation(
			r,
			fmt.Sprintf(
				"target package not found: %s - pass create to create it",
				targetPackagePath,
			),
		)
	}

	newNames := make(map[string]string)

	for _, entry := range entries {
		if previous, taken := newNames[entry.newName]; taken {
			return failValidation(
				r,
				fmt.Sprintf(
					"%s and %s both move as %s",
					previous,
					entry.symbol,
					entry.newName,
				),
			)
		}

		newNames[entry.newName] = entry.symbol

		if target != nil {
			if f := checkScopeCollision(target, entry.newName); f != nil {
				return failValidation(r, f.Error())
			}
		}
	}

	if target != nil && importsTransitively(target, packagePath) {
		return failValidation(
			r,
			fmt.Sprintf(
				"move would create an import cycle: %s imports %s",
				targetPackagePath,
				packagePath,
			),
		)
	}

	excluded := make(map[token.Pos]bool)

	for _, entry := range entries {
		excluded[entry.object.Pos()] = true
	}

	for _, entry := range entries {
		dependencies := moveDependencies(p, excluded, entry.node)

		if len(dependencies) > 0 {
			return failValidation(
				r,
				fmt.Sprintf(
					"%s references package-local symbols: %s",
					entry.symbol,
					strings.Join(dependencies, ", "),
				),
			)
		}

		for _, c := range entry.carried {
			importPath, f := strconv.Unquote(c.Path.Value)

			if f != nil {
				continue
			}

			if importPath == targetPackagePath {
				return failValidation(
					r,
					fmt.Sprintf(
						"%s references the target package",
						entry.symbol,
					),
				)
			}

			carriedPackage := findPackage(all, importPath)

			if target != nil && carriedPackage != nil &&
				importsTransitively(carriedPackage, targetPackagePath) {
				return failValidation(
					r,
					fmt.Sprintf(
						"move would create an import cycle through %s",
						importPath,
					),
				)
			}
		}
	}

	moveDirectory, e := targetDirectory(p, target, targetPackagePath)

	if e != nil {
		return failValidation(r, e.Error())
	}

	targetPackageName := path.Base(targetPackagePath)

	if target != nil {
		targetPackageName = target.Types.Name()
	}

	var external []qualifiedReference
	renames := make(map[*ast.Ident]string)

	for _, entry := range entries {
		for _, f := range resolve.FindAllReferences(all, entry.object) {
			if insideMoved(entries, f.Ident.Pos()) {
				if entry.flipped {
					renames[f.Ident] = entry.newName
				}

				continue
			}

			external = append(
				external,
				qualifiedReference{reference: f, newName: entry.newName},
			)
		}
	}

	qualifications, blocked := planQualifications(
		set,
		external,
		packagePath,
		targetPackagePath,
		targetPackageName,
	)

	if qualifications == nil {
		return failValidation(
			r,
			fmt.Sprintf("no available import name in %s", blocked),
		)
	}

	return executeMove(
		r,
		&movePlan{
			set:               set,
			entries:           entries,
			qualifications:    qualifications,
			renames:           renames,
			packagePath:       packagePath,
			targetPackagePath: targetPackagePath,
			targetPackageName: targetPackageName,
			moveDirectory:     moveDirectory,
			createTarget:      target == nil,
		},
	)
}
