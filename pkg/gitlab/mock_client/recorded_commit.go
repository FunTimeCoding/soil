package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

type RecordedCommit struct {
	Branch  string
	Message string
	Actions []*gitlab.CommitActionOptions
}
