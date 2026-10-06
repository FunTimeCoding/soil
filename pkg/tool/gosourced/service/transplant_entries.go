package service

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"sort"
)

func transplantEntries(
	d *decoration.Set,
	source *packages.Package,
	entries []*relocation.Entry,
) []dst.Decl {
	ordered := append([]*relocation.Entry{}, entries...)
	sort.Slice(
		ordered,
		func(
			i int,
			j int,
		) bool {
			return ordered[i].Object.Pos() < ordered[j].Object.Pos()
		},
	)
	seen := make(map[ast.Node]bool)
	var result []dst.Decl
	groupIndex := make(map[token.Token]int)
	groups := make(map[token.Token][]*relocation.TransplantSpec)

	for _, entry := range ordered {
		if entry.Spec == nil {
			if seen[entry.Declaration] {
				continue
			}

			seen[entry.Declaration] = true
			declaration := d.DecoratedNode(source, entry.Declaration).(dst.Decl)
			declaration.Decorations().Before = dst.EmptyLine
			result = append(result, declaration)

			continue
		}

		if seen[entry.Spec] {
			continue
		}

		seen[entry.Spec] = true
		g := entry.Declaration.(*ast.GenDecl)
		declaration := d.DecoratedNode(source, entry.Declaration).(*dst.GenDecl)
		spec := d.DecoratedNode(source, entry.Spec).(dst.Spec)
		single := len(g.Specs) == 1

		if g.Tok == token.TYPE {
			result = append(result, transplantSingle(declaration, spec, single))

			continue
		}

		if _, exists := groupIndex[g.Tok]; !exists {
			groupIndex[g.Tok] = len(result)
			result = append(result, nil)
		}

		groups[g.Tok] = append(
			groups[g.Tok],
			relocation.NewTransplantSpec(declaration, spec, single))
	}

	for tok, parts := range groups {
		i := groupIndex[tok]

		if len(parts) == 1 {
			result[i] = transplantSingle(
				parts[0].Declaration,
				parts[0].Spec,
				parts[0].Single,
			)

			continue
		}

		parent := parts[0].Declaration
		whole := len(parts) == len(parent.Specs)

		for _, part := range parts {
			if part.Declaration != parent {
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
			counts[part.Declaration]++
		}

		merged := &dst.GenDecl{Tok: tok, Lparen: true, Rparen: true}
		merged.Decs.Before = dst.EmptyLine
		carried := make(map[*dst.GenDecl]bool)

		for _, part := range parts {
			absorbed := counts[part.Declaration] ==
				len(part.Declaration.Specs)

			if absorbed && !carried[part.Declaration] {
				carried[part.Declaration] = true
				part.Spec.Decorations().Start.Prepend(
					part.Declaration.Decs.Start.All()...,
				)
			}

			if part.Spec.Decorations().Before == dst.None {
				part.Spec.Decorations().Before = dst.NewLine
			}

			merged.Specs = append(merged.Specs, part.Spec)
		}

		result[i] = merged
	}

	return result
}
