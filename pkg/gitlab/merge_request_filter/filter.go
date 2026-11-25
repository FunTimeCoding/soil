package merge_request_filter

import "time"

type Filter struct {
	Project      int64
	Limit        int
	Offset       int
	State        string
	OrderBy      string
	UpdatedAfter *time.Time
	Author       string
	Reviewer     string
	Assignee     string
	SourceBranch string
	TargetBranch string
	Search       string
	Draft        *bool
}
