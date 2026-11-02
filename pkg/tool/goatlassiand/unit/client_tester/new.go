package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"testing"
)

func New(
	t *testing.T,
	j face.JiraClient,
) *client.Client {
	t.Helper()
	result, _ := NewReported(t, j)

	return result
}
