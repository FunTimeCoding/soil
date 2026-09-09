package merge_request_detail

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Detail struct {
	Project    int64
	Identifier int64
	Title      string
	State      string
	Link       string
	Create     *time.Time
	Raw        *gitlab.MergeRequest
}
