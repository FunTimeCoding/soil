package gocheck

import (
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/goreach"
)

func reached(o *option.Check) (*reach.Reach, string, error) {
	r, edits, files, e := goreach.Reach(o.Trees, o.Root, o.Transcript)

	if e != nil {
		return nil, "", e
	}

	return r, goreach.Announcement(r, edits, files), nil
}
