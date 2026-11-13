package gocheck

import (
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func fromResponse(r client.CheckResponse) *check_result.Result {
	result := check_result.New(r.Root)
	result.Announcement = r.Announcement
	result.Stale = r.Stale
	result.Unchanged = r.Unchanged
	result.Checked = r.Checked

	for _, s := range r.Sections {
		result.Sections = append(result.Sections, fromSection(r.Root, s))
	}

	for _, h := range r.Changes {
		result.Changes = append(
			result.Changes,
			layer_change.New(h.Path, h.Destination, h.Original, h.Current),
		)
	}

	return result
}
