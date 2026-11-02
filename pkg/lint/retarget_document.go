package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	"github.com/funtimecoding/soil/pkg/markup/front_matter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func RetargetDocument(
	r *resolver.Resolver,
	path string,
	content string,
	moves *move_set.Set,
) (string, []*retarget_change.Change) {
	through := retarget.Through(r, moves.Paths)
	declared := pointer.NewDeclared()
	skip := 0
	baseLine := 0

	if f, found := front_matter.Extract(content); found {
		skip = f.Lines
		baseLine = f.Line(constant.BaseKey)
		var d Declaration

		if f.Decode(&d) == nil {
			for _, value := range declaredValues(d.Base) {
				if through.BaseExists(value) {
					declared.Bases = append(declared.Bases, value)
				}
			}
		}
	}

	lines := strings.Split(content, stringsConstant.Unix)
	var changes []*retarget_change.Change

	for i, line := range lines {
		number := i + 1

		if number <= skip {
			if baseBlock(lines, baseLine, number) {
				lines[i] = retarget.Values(moves.Paths, declared.Bases, line)
			}

			continue
		}

		rewritten, found := retarget.Line(
			through,
			path,
			declared,
			moves,
			line,
			number,
		)
		lines[i] = rewritten
		changes = append(changes, found...)
	}

	return strings.Join(lines, stringsConstant.Unix), changes
}
