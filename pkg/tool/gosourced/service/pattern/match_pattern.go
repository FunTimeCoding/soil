package pattern

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern/literals"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern/matcher"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
)

func (s *Service) MatchPattern(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
	pattern string,
) (*output.Results, *result.Match, error) {
	r := output.NewResultsWithDirectory(directory)
	specification, e := match.Parse(pattern)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil, nil
	}

	all, set, f := s.workspace.CensusPackages(directory, packagePath)

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
	qualifiers, _ := qualifierPaths(
		specification,
		symbol,
		anchorPackage(declaration, receiver),
		lookup.PackageNames(all),
	)
	var patternNode ast.Node = specification.Statement
	call := expressionCall(specification, symbol)

	if call != nil {
		patternNode = call
	}

	contents := map[string][]byte{}
	total := 0
	matched := 0
	var entries []*pattern_site.Entry

	for _, reference := range lookup.ObjectReferences(all, isAnchor) {
		file := lookup.SyntaxFile(reference.Package, reference.Ident.Pos())

		if file == nil {
			continue
		}

		total++
		path, _ := astutil.PathEnclosingInterval(
			file,
			reference.Ident.Pos(),
			reference.Ident.End(),
		)
		node, anchor := siteSelection(path, call != nil)
		unification := matcher.New(
			specification,
			qualifiers,
			reference.Package.TypesInfo,
			reference.Package.Types,
			set,
			symbol,
			isAnchor,
		)

		if unification.Unify(patternNode, node) {
			matched++

			continue
		}

		entry, i := s.siteEntry(
			directory,
			set,
			contents,
			node,
			anchor,
			reference,
		)

		if i != nil {
			return nil, nil, i
		}

		entries = append(entries, entry)
	}

	unmatched := literals.GroupEntries(entries)

	return r, result.NewMatch(symbol, pattern, total, matched, unmatched), nil
}
