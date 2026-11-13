package pattern

import (
	"fmt"
	"github.com/dave/dst"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern/literals"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern/matcher"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
)

func (s *Service) applyPatternThrough(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
	pattern string,
	replacement string,
	partial bool,
	out *sink.Sink,
) (*output.Results, *result.Apply, error) {
	r := output.NewResultsWithDirectory(directory)
	patternSpec, replacementSpec, e := parseApplyPatterns(
		pattern,
		replacement,
		symbol,
	)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil, nil
	}

	all, set, f := load.Packages(directory, goModConstant.AllPackages)

	if f != nil {
		return nil, nil, f
	}

	declaration, _, g := lookup.FindDeclaration(
		all,
		packagePath,
		symbol,
		receiver,
	)

	if g != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, g.Error(), "", false),
		)

		return r, nil, nil
	}

	targets, h := anchorSet(declaration, receiver)

	if h != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, h.Error(), "", false),
		)

		return r, nil, nil
	}

	isAnchor := func(o types.Object) bool {
		for _, target := range targets {
			if lookup.SameObject(o, target) {
				return true
			}
		}

		return false
	}
	anchor := anchorPackage(declaration, receiver)
	qualifiers, i := resolveQualifiers(
		replacementSpec,
		symbol,
		anchor,
		lookup.PackageNames(all),
	)

	if i != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, i.Error(), "", false),
		)

		return r, nil, nil
	}

	matching, _ := qualifierPaths(
		patternSpec,
		symbol,
		anchor,
		lookup.PackageNames(all),
	)
	var patternNode ast.Node = patternSpec.Statement
	call := expressionCall(patternSpec, symbol)

	if call != nil {
		patternNode = call

		if _, okay := replacementSpec.Statement.(*ast.ExprStmt); !okay {
			r.AddConcern(
				concern.NewFile(
					constant.ConcernValidation,
					"the pattern is a call, so the replacement must be an expression",
					"",
					false,
				),
			)

			return r, nil, nil
		}
	}

	report := result.NewApply(symbol, pattern, replacement)
	decorations := decoration.NewSet()
	contents := map[string][]byte{}
	seen := map[token.Pos]bool{}
	var plans []pattern_site.ApplyPlan
	var planned []*pattern_site.Entry
	var unmatched []*pattern_site.Entry
	var refused []*pattern_site.Entry

	for _, reference := range lookup.ObjectReferences(all, isAnchor) {
		file := lookup.SyntaxFile(reference.Package, reference.Ident.Pos())

		if file == nil {
			continue
		}

		report.Total++
		path, _ := astutil.PathEnclosingInterval(
			file,
			reference.Ident.Pos(),
			reference.Ident.End(),
		)
		node, shape := siteSelection(path, call != nil)
		entry, j := s.siteEntry(
			directory,
			set,
			contents,
			node,
			shape,
			reference,
		)

		if j != nil {
			return nil, nil, j
		}

		unification := matcher.New(
			patternSpec,
			matching,
			reference.Package.TypesInfo,
			reference.Package.Types,
			set,
			symbol,
			isAnchor,
		)

		if !unification.Unify(patternNode, node) {
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

		if parent == nil && call == nil {
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

		if old == nil || decoration.HasDecorations(old) {
			refused = append(refused, entry)

			continue
		}

		plans = append(
			plans,
			pattern_site.ApplyPlan{
				Package:   reference.Package,
				Statement: node,
				Parent:    parent,
				Bindings:  unification.Bindings(),
				Anchor:    anchorNode(path),
			},
		)
		planned = append(planned, entry)
	}

	nested := nestedPlans(plans)
	var kept []pattern_site.ApplyPlan

	for i, plan := range plans {
		if nested[i] {
			refused = append(refused, planned[i])

			continue
		}

		kept = append(kept, plan)
	}

	plans = kept
	report.Unmatched = literals.GroupEntries(unmatched)
	report.Refused = literals.GroupEntries(refused)

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
			plan.Package,
			plan.Bindings,
			plan.Anchor,
		)

		if l != nil {
			return nil, nil, l
		}

		if call != nil {
			if m := swapExpression(
				decorations.Files[set.Position(plan.Statement.Pos()).Filename],
				decorations.DecoratedNode(plan.Package, plan.Statement),
				statement.(*dst.ExprStmt).X,
			); m != nil {
				return nil, nil, m
			}

			continue
		}

		if m := swapStatement(
			decorations.DecoratedNode(plan.Package, plan.Parent),
			decorations.DecoratedNode(plan.Package, plan.Statement),
			statement,
		); m != nil {
			return nil, nil, m
		}
	}

	changed, n := writeChanged(decorations, resolve.NewNames(all), out)

	if n != nil {
		return nil, nil, n
	}

	for _, plan := range plans {
		if changed[set.Position(plan.Statement.Pos()).Filename] {
			report.Rewritten++
		}
	}

	if len(plans) > 0 && report.Rewritten == 0 {
		report.Refusal = "the replacement changes no file - it prints as the original at every matched site"

		return r, report, nil
	}

	report.Applied = true

	return r, report, nil
}
