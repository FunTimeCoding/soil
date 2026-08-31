package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/match"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
)

func (s *Service) ApplyPattern(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
	pattern string,
	replacement string,
	partial bool,
	dryRun bool,
) (*output.Results, *result.Apply, error) {
	r := output.NewResultsWithDirectory(directory)
	patternSpec, replacementSpec, e := parseApplyPatterns(
		pattern,
		replacement,
		symbol,
	)

	if e != nil {
		r.AddConcern(concern.NewFile("validation", e.Error(), "", false))

		return r, nil, nil
	}

	all, set, f := resolve.LoadPackages(directory, "./...")

	if f != nil {
		return nil, nil, f
	}

	declaration, _, g := findDeclaration(all, packagePath, symbol, receiver)

	if g != nil {
		r.AddConcern(concern.NewFile("validation", g.Error(), "", false))

		return r, nil, nil
	}

	targets, h := anchorSet(declaration, receiver)

	if h != nil {
		r.AddConcern(concern.NewFile("validation", h.Error(), "", false))

		return r, nil, nil
	}

	isAnchor := func(o types.Object) bool {
		for _, target := range targets {
			if sameObject(o, target) {
				return true
			}
		}

		return false
	}
	qualifiers, i := resolveQualifiers(
		replacementSpec,
		symbol,
		packageNames(all),
	)

	if i != nil {
		r.AddConcern(concern.NewFile("validation", i.Error(), "", false))

		return r, nil, nil
	}

	report := result.NewApply(symbol, pattern, replacement)
	decorations := decoration.NewSet()
	contents := map[string][]byte{}
	seen := map[token.Pos]bool{}
	var plans []applyPlan
	var unmatched []*siteEntry
	var refused []*siteEntry

	for _, reference := range objectReferences(all, isAnchor) {
		file := syntaxFileAt(reference.Package, reference.Ident.Pos())

		if file == nil {
			continue
		}

		report.Total++
		path, _ := astutil.PathEnclosingInterval(
			file,
			reference.Ident.Pos(),
			reference.Ident.End(),
		)
		node := siteNode(path)
		entry, j := s.siteEntryFor(
			directory,
			set,
			contents,
			node,
			anchorNode(path),
			reference,
		)

		if j != nil {
			return nil, nil, j
		}

		unification := match.New(
			patternSpec,
			reference.Package.TypesInfo,
			reference.Package.Types,
			set,
			symbol,
			isAnchor,
		)

		if !unification.Unify(patternSpec.Statement, node) {
			unmatched = append(unmatched, entry)

			continue
		}

		report.Matched++

		if seen[node.Pos()] {
			refused = append(refused, entry)

			continue
		}

		seen[node.Pos()] = true
		parent := statementParent(path, node)

		if parent == nil {
			refused = append(refused, entry)

			continue
		}

		if _, k := decorations.DecorateFile(
			set,
			reference.Package,
			file,
		); k != nil {
			return nil, nil, k
		}

		old := decorations.DecoratedNode(reference.Package, node)

		if old == nil || hasDecorations(old) {
			refused = append(refused, entry)

			continue
		}

		plans = append(
			plans,
			applyPlan{
				p:        reference.Package,
				statement:     node,
				parent:   parent,
				bindings: unification.Bindings(),
				anchor:   anchorNode(path),
			},
		)
	}

	report.Unmatched = groupEntries(unmatched)
	report.Refused = groupEntries(refused)

	if !partial && len(unmatched)+len(refused) > 0 {
		report.Refusal = fmt.Sprintf(
			"population not uniform: %d unmatched, %d refused - pass partial to rewrite the matched sites",
			len(unmatched),
			len(refused),
		)

		return r, report, nil
	}

	for _, plan := range plans {
		statement, l := buildReplacement(
			replacementSpec,
			replacement,
			symbol,
			qualifiers,
			decorations,
			plan.p,
			plan.bindings,
			plan.anchor,
		)

		if l != nil {
			return nil, nil, l
		}

		if m := swapStatement(
			decorations.DecoratedNode(plan.p, plan.parent),
			decorations.DecoratedNode(plan.p, plan.statement),
			statement,
		); m != nil {
			return nil, nil, m
		}

		report.Rewritten++
	}

	resolver := resolve.NewNames(all)

	for filename, file := range decorations.Files {
		if n := restoreDecoratedFile(
			resolver,
			decorations.PackagePaths[file],
			map[string]string{},
			file,
			filename,
			dryRun,
		); n != nil {
			return nil, nil, n
		}
	}

	report.Applied = true

	return r, report, nil
}
