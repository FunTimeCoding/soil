package issue

import "gitlab.com/gitlab-org/api/client-go/v3"

type Issue struct {
	Project    int64
	Identifier int64
	Title      string
	State      string
	Link       string
	Raw        *gitlab.Issue
}
