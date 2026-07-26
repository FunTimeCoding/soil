package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

func renderEntries(
	plan *movePlan,
	entries []*moveEntry,
	sources map[string][]byte,
	edits []*spliceEdit,
) string {
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
	var blocks []string
	groupIndex := make(map[token.Token]int)
	groups := make(map[token.Token][]*renderPart)

	for _, entry := range ordered {
		filename := plan.set.Position(entry.file.Pos()).Filename
		content := sources[filename]
		tokenFile := plan.set.File(entry.file.Pos())

		if entry.spec == nil {
			if seen[entry.declaration] {
				continue
			}

			seen[entry.declaration] = true
			start := entry.declaration.Pos()

			if d := declarationDocument(entry.declaration); d != nil {
				start = d.Pos()
			}

			end := trailingCommentEnd(
				plan.set,
				entry.file,
				entry.declaration.End(),
			)
			blocks = append(
				blocks,
				applySpliceEdits(content, tokenFile, start, end, edits),
			)

			continue
		}

		if seen[entry.spec] {
			continue
		}

		seen[entry.spec] = true
		document := ""

		if d := specDocument(entry.declaration, entry.spec); d != nil {
			document = fmt.Sprintf(
				"%s\n",
				content[tokenFile.Offset(d.Pos()):tokenFile.Offset(d.End())],
			)
		}

		body := applySpliceEdits(
			content,
			tokenFile,
			entry.spec.Pos(),
			specSliceEnd(entry.spec),
			edits,
		)

		if _, okay := entry.spec.(*ast.TypeSpec); okay {
			blocks = append(
				blocks,
				fmt.Sprintf("%stype %s", document, body),
			)

			continue
		}

		g := entry.declaration.(*ast.GenDecl)

		if _, exists := groupIndex[g.Tok]; !exists {
			groupIndex[g.Tok] = len(blocks)
			blocks = append(blocks, "")
		}

		groups[g.Tok] = append(
			groups[g.Tok],
			&renderPart{document: document, body: body},
		)
	}

	for tok, parts := range groups {
		i := groupIndex[tok]

		if len(parts) == 1 {
			blocks[i] = fmt.Sprintf(
				"%s%s %s",
				parts[0].document,
				tok,
				parts[0].body,
			)

			continue
		}

		var lines []string

		for _, p := range parts {
			lines = append(lines, join.Empty(p.document, p.body))
		}

		blocks[i] = fmt.Sprintf(
			"%s (\n%s\n)",
			tok,
			strings.Join(lines, "\n"),
		)
	}

	return strings.Join(blocks, "\n\n")
}
