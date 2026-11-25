package argument

type ListMergeRequests struct {
	Project      string `json:"project"`
	State        string `json:"state"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	OrderBy      string `json:"order_by"`
	UpdatedAfter string `json:"updated_after"`
	Author       string `json:"author"`
	Reviewer     string `json:"reviewer"`
	Assignee     string `json:"assignee"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Search       string `json:"search"`
	Draft        *bool  `json:"draft"`
}
