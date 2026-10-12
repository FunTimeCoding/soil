package service_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"testing"
)

func Service(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	s := store.New(lite.NewMemory())
	t.Cleanup(s.Close)

	return service.New(c, s)
}
