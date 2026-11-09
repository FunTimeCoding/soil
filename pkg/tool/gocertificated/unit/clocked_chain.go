package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"testing"
)

func clockedChain(t *testing.T) (*service.Service, *record.Record) {
	t.Helper()
	v, _ := service_tester.NewWith(t, windowNow)
	root, e := v.CreateAuthority(server_tester.RootBody())
	assert.FatalOnError(t, e)
	_, f := v.CreateAuthority(server_tester.ClusterBody())
	assert.FatalOnError(t, f)

	return v, root
}
