package service

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"sort"
)

func transplantEntries(
	d *decoration.Set,
	source *packages.Package,
	entries []*moveEntry,
) []dst.Decl {
	ordered := append([]*moveEntry{}, entries...)
	sort.Slice(
		ordered,
		func(
			i int,
			j int,
		) bool {
			return ordered[i].object.Pos() < ordered[j].object.Pos()
		},
	)
	seen := make(map[ast.Node]bool)
	var result []dst.Decl
	groupIndex := make(map[token.Token]int)
	groups := make(map[token.Token][]*transplantSpec)

	for _, entry := range ordered {
		if entry.spec == nil {
			if seen[entry.declaration] {
				continue
			}

			seen[entry.declaration] = true
			declaration := d.DecoratedNode(source, entry.declaration).(dst.Decl)
			declaration.Decorations().Before = dst.EmptyLine
			result = append(result, declaration)

			continue
		}

		if seen[entry.spec] {
			continue
		}

		seen[entry.spec] = true
		g := entry.declaration.(*ast.GenDecl)
		declaration := d.DecoratedNode(source, entry.declaration).(*dst.GenDecl)
		spec := d.DecoratedNode(source, entry.spec).(dst.Spec)
		single := len(g.Specs) == 1

		if g.Tok == token.TYPE {
			result = append(
				result,
				transplantSingle(declaration, spec, single),
			)

			continue
		}

		if _, exists := groupIndex[g.Tok]; !exists {
			groupIndex[g.Tok] = len(result)
			result = append(result, nil)
		}

		groups[g.Tok] = append(
			groups[g.Tok],
			&transplantSpec{
				declaration: declaration,
				spec:        spec,
				single:      single,
			},
		)
	}

	for tok, parts := range groups {
		i := groupIndex[tok]

		if len(parts) == 1 {
			result[i] = transplantSingle(
				parts[0].declaration,
				parts[0].spec,
				parts[0].single,
			)

			continue
		}

		parent := parts[0].declaration
		whole := len(parts) == len(parent.Specs)

		for _, part := range parts {
			if part.declaration != parent {
				whole = false
			}
		}

		if whole {
			clone := &dst.GenDecl{
				Tok:    parent.Tok,
				Lparen: parent.Lparen,
				Rparen: parent.Rparen,
				Specs:  append([]dst.Spec{}, parent.Specs...),
			}
			clone.Decs = parent.Decs
			clone.Decs.Before = dst.EmptyLine
			result[i] = clone

			continue
		}

		counts := make(map[*dst.GenDecl]int)

		for _, part := range parts {
			counts[part.declaration]++
		}

		merged := &dst.GenDecl{
			Tok:    tok,
			Lparen: true,
			Rparen: true,
		}
		merged.Decs.Before = dst.EmptyLine
		carried := make(map[*dst.GenDecl]bool)

		for _, part := range parts {
			absorbed := counts[part.declaration] ==
				len(part.declaration.Specs)

			if absorbed && !carried[part.declaration] {
				carried[part.declaration] = true
				part.spec.Decorations().Start.Prepend(
					part.declaration.Decs.Start.All()...,
				)
			}

			if part.spec.Decorations().Before == dst.None {
				part.spec.Decorations().Before = dst.NewLine
			}

			merged.Specs = append(merged.Specs, part.spec)
		}

		result[i] = merged
	}

	return result
}
