package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
)

func (o *Tester) Chain() *service.Service {
	o.t.Helper()
	v := o.Service()
	_, e := v.CreateAuthority(server_tester.RootBody())
	assert.FatalOnError(o.t, e)
	_, f := v.CreateAuthority(server_tester.ClusterBody())
	assert.FatalOnError(o.t, f)

	return v
}
