package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"testing"
)

func NewChain(t *testing.T) *service.Service {
	t.Helper()
	v := NewService(t)
	_, e := v.CreateAuthority(server_tester.RootBody())
	assert.Nil(t, e)
	_, f := v.CreateAuthority(server_tester.ClusterBody())
	assert.Nil(t, f)

	return v
}
