package recorded_commit

import "gitlab.com/gitlab-org/api/client-go/v3"

type Commit struct {
	Branch  string
	Message string
	Actions []*gitlab.CommitActionOptions
}
