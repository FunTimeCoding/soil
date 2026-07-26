package service

import (
	"go/token"
	"sort"
)

func applySpliceEdits(
	content []byte,
	file *token.File,
	start token.Pos,
	end token.Pos,
	edits []*spliceEdit,
) string {
	var scoped []*spliceEdit

	for _, edit := range edits {
		if edit.start >= start && edit.end <= end {
			scoped = append(scoped, edit)
		}
	}

	sort.Slice(
		scoped,
		func(
			i int,
			j int,
		) bool {
			return scoped[i].start > scoped[j].start
		},
	)
	segment := append(
		[]byte{},
		content[file.Offset(start):file.Offset(end)]...,
	)

	for _, edit := range scoped {
		s := file.Offset(edit.start) - file.Offset(start)
		t := file.Offset(edit.end) - file.Offset(start)
		segment = append(
			segment[:s],
			append([]byte(edit.replacement), segment[t:]...)...,
		)
	}

	return string(segment)
}
