package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"path"
	"path/filepath"
	"sort"
)

func (s *Service) RenamePackage(
	directory string,
	packagePath string,
	newName string,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)

	if !token.IsIdentifier(newName) {
		return failValidation(
			r,
			fmt.Sprintf("not a valid package name: %s", newName),
		)
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

	if len(p.GoFiles) == 0 {
		return failValidation(
			r,
			fmt.Sprintf("package has no Go files: %s", packagePath),
		)
	}

	oldName := p.Types.Name()

	if oldName == newName {
		return failValidation(
			r,
			fmt.Sprintf("package is already named %s", newName),
		)
	}

	targetPackagePath := path.Join(path.Dir(packagePath), newName)

	if targetPackagePath == packagePath {
		return failValidation(
			r,
			fmt.Sprintf("import path would not change: %s", packagePath),
		)
	}

	moveDirectory, e := targetDirectory(p, nil, targetPackagePath)

	if e != nil {
		return failValidation(r, e.Error())
	}

	if _, f := os.Stat(moveDirectory); f == nil {
		return failValidation(
			r,
			fmt.Sprintf(
				"target directory already exists: %s",
				moveDirectory,
			),
		)
	}

	var qualifiers []resolve.Reference

	for _, loaded := range all {
		for ident, o := range loaded.TypesInfo.Uses {
			packageName, okay := o.(*types.PkgName)

			if !okay || packageName.Imported().Path() != packagePath {
				continue
			}

			if packageName.Name() != oldName {
				continue
			}

			scope := loaded.Types.Scope().Innermost(ident.Pos())

			if scope != nil {
				_, shadow := scope.LookupParent(newName, ident.Pos())

				if shadow != nil &&
					shadow.Parent() != types.Universe {
					position := set.Position(ident.Pos())

					return failValidation(
						r,
						fmt.Sprintf(
							"%s is taken at %s:%d",
							newName,
							position.Filename,
							position.Line,
						),
					)
				}
			}

			qualifiers = append(
				qualifiers,
				resolve.Reference{Ident: ident, Package: loaded},
			)
		}
	}

	modified := make(map[string]*ast.File)
	rewriteImportPaths(r, set, all, packagePath, targetPackagePath, modified)

	for _, f := range qualifiers {
		position := set.Position(f.Ident.Pos())
		f.Ident.Name = newName
		filename := position.Filename
		file := findSyntaxFile(set, f.Package, filename)

		if file != nil {
			modified[filename] = file
		}

		r.AddConcern(
			concern.NewLine(
				"renamed",
				fmt.Sprintf("%s → %s", oldName, newName),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	sourceDirectory := filepath.Dir(p.GoFiles[0])
	testName := fmt.Sprintf("%s_test", oldName)

	for _, loaded := range all {
		for _, file := range loaded.Syntax {
			filename := set.Position(file.Pos()).Filename

			if filepath.Dir(filename) != sourceDirectory {
				continue
			}

			if file.Name.Name == oldName {
				file.Name.Name = newName
				modified[filename] = file
			}

			if file.Name.Name == testName {
				file.Name.Name = fmt.Sprintf("%s_test", newName)
				modified[filename] = file
			}
		}
	}

	var filenames []string

	for filename := range modified {
		filenames = append(filenames, filename)
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		e = writeFile(set, modified[filename], filename)

		if e != nil {
			return nil, e
		}
	}

	e = os.Rename(sourceDirectory, moveDirectory)

	if e != nil {
		return nil, e
	}

	r.AddConcern(
		concern.NewFile(
			"moved",
			fmt.Sprintf("%s → %s", packagePath, targetPackagePath),
			moveDirectory,
			true,
		),
	)

	return r, nil
}
