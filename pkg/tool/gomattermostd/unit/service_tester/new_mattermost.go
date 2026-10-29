package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"testing"
)

func NewMattermost(
	t *testing.T,
	me string,
) *mock_client.Client {
	t.Helper()
	result := mock_client.New(me)
	t.Cleanup(result.Close)

	return result
}
