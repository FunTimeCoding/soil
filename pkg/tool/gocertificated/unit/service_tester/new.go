package service_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"testing"
	"time"
)

func New(t *testing.T) (*service.Service, *mock_client.Client) {
	t.Helper()

	return NewWith(t, time.Now)
}
