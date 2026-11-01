package gosource

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosource/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"strconv"
)

func DiffLines(r client.DiffResponse) []string {
	counts := []string{constant.CountsLabel}

	for _, c := range r.Counts {
		counts = append(counts, c.Class, strconv.Itoa(c.Count))
	}

	result := []string{join.Space(counts...)}

	for _, m := range r.Modules {
		to := m.To

		if to == "" {
			to = constant.StagedLabel
		}

		result = append(
			result,
			join.Space(constant.ModuleLabel, m.Name, m.From, to),
		)
	}

	for _, file := range r.FromUnparsed {
		result = append(
			result,
			join.Space(constant.UnparsedLabel, constant.FromSide, file),
		)
	}

	for _, file := range r.ToUnparsed {
		result = append(
			result,
			join.Space(constant.UnparsedLabel, constant.ToSide, file),
		)
	}

	for _, e := range r.Entries {
		result = append(result, entryLine(e))
	}

	return result
}
