package service

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"strings"
)

func retargetMemory(
	r *resolver.Resolver,
	moves *move_set.Set,
	m *record.Memory,
) (string, string, []*retarget_change.Change) {
	base := m.Metadata[constant.BaseKey]
	bases := reference.Bases(base)
	declared := pointer.NewDeclared()

	for _, b := range bases {
		if r.BaseExists(b) {
			declared.Bases = append(declared.Bases, b)
		}
	}

	lines := strings.Split(m.Content, stringsConstant.Unix)
	var changes []*retarget_change.Change

	for i, line := range lines {
		rewritten, found := retarget.Line(r, "", declared, moves, line, i+1)
		lines[i] = rewritten
		changes = append(changes, found...)
	}

	movedBase := retarget.Values(moves.Paths, bases, base)

	if movedBase != base {
		changes = append(changes, retarget_change.New(0, base, movedBase))
	}

	return strings.Join(lines, stringsConstant.Unix), movedBase, changes
}
