package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client/mock_client"
	"testing"
)

func NewParser(t *testing.T) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	t.Cleanup(result.Close)

	return result
}
