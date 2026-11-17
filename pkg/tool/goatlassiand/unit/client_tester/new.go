package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"testing"
)

func New(
	t *testing.T,
	j *mock_client.Client,
) *client.Client {
	t.Helper()
	result, _ := NewReported(t, j)

	return result
}
