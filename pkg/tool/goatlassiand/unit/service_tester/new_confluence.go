package service_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	"testing"
)

func NewConfluence(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
