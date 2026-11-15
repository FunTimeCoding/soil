package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"testing"
)

func ClockedChain(t *testing.T) (*service.Service, *record.Record) {
	t.Helper()
	v, _ := NewWith(t, clock_tester.Window)
	root, e := v.CreateAuthority(server_tester.RootBody())
	assert.FatalOnError(t, e)
	_, f := v.CreateAuthority(server_tester.ClusterBody())
	assert.FatalOnError(t, f)

	return v, root
}
