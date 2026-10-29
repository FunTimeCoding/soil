package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

type Client struct {
	Failure            error
	Closed             int
	projects           []response.Project
	issues             []*response.Issue
	issueSearchProject string
	eventSearchProject string
}
