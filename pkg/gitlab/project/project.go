package project

import "gitlab.com/gitlab-org/api/client-go/v3"

type Project struct {
	Identifier int64
	Namespace  string
	Name       string
	Link       string
	Raw        *gitlab.Project
}
