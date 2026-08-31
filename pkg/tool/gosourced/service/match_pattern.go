package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/match"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
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
	contents := map[string][]byte{}
	total := 0
	matched := 0
	var entries []*siteEntry

	for _, reference := range objectReferences(all, isAnchor) {
		file := syntaxFileAt(reference.Package, reference.Ident.Pos())

		if file == nil {
			continue
		}

		total++
		path, _ := astutil.PathEnclosingInterval(
			file,
			reference.Ident.Pos(),
			reference.Ident.End(),
		)
		node := siteNode(path)
		unification := match.New(
			specification,
			reference.Package.TypesInfo,
			reference.Package.Types,
			set,
			symbol,
			isAnchor,
		)

		if unification.Unify(specification.Statement, node) {
			matched++

			continue
		}

		entry, i := s.siteEntryFor(
			directory,
			set,
			contents,
			node,
			anchorNode(path),
			reference,
		)

		if i != nil {
			return nil, nil, i
		}

		entries = append(entries, entry)
	}

	unmatched := groupEntries(entries)

	return r, result.NewMatch(symbol, pattern, total, matched, unmatched), nil
}
