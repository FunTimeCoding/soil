package transcript

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/memory_payload"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile"
	"slices"
)

func textEntries(
	body string,
	lines []string,
	section string,
) ([]memory_payload.Entry, bool) {
	headings := profile.Headings(body, section)

	if len(headings) == 0 {
		return nil, !slices.Contains(lines, section)
	}

	result := make([]memory_payload.Entry, 0, len(headings))

	for _, h := range headings {
		result = append(result, *memory_payload.NewEntry(h.Identifier, h.Name))
	}

	return result, true
}
