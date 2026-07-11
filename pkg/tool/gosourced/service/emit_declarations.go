package service

import (
	"go/ast"
	"go/token"
	"sort"
)

func emitDeclarations(entries []*moveEntry) []ast.Decl {
	ordered := append([]*moveEntry{}, entries...)
	sort.Slice(
		ordered,
		func(i int, j int) bool {
			return ordered[i].object.Pos() < ordered[j].object.Pos()
		},
	)
	var result []ast.Decl
	seen := make(map[ast.Node]bool)
	groups := make(map[token.Token][]ast.Spec)

	for _, entry := range ordered {
		if entry.spec == nil {
			if !seen[entry.declaration] {
				seen[entry.declaration] = true
				result = append(result, entry.declaration)
			}

			continue
		}

		if seen[entry.spec] {
			continue
		}

		seen[entry.spec] = true
		g := entry.declaration.(*ast.GenDecl)

		if _, okay := entry.spec.(*ast.TypeSpec); okay {
			result = append(
				result,
				&ast.GenDecl{Tok: g.Tok, Specs: []ast.Spec{entry.spec}},
			)

			continue
		}

		if len(groups[g.Tok]) == 0 {
			placeholder := &ast.GenDecl{Tok: g.Tok}
			result = append(result, placeholder)
		}

		groups[g.Tok] = append(groups[g.Tok], entry.spec)
	}

	for i, d := range result {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok == token.TYPE || len(g.Specs) > 0 {
			continue
		}

		specs := groups[g.Tok]
		merged := &ast.GenDecl{Tok: g.Tok, Specs: specs}

		if len(specs) > 1 {
			merged.Lparen = 1
		}

		result[i] = merged
	}

	return result
}
