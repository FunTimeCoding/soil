package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

type Client struct {
	projects           []response.Project
	issues             []*response.Issue
	issueSearchProject string
	eventSearchProject string
}
