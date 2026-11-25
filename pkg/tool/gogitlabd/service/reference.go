package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func reference(d *merge_request_detail.Detail) string {
	if r := d.Raw.References; r != nil && r.Full != "" {
		return r.Full
	}

	return fmt.Sprintf("%d!%d", d.Project, d.Identifier)
}
