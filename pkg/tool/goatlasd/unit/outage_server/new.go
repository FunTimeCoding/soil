package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	gooutpostdConstant "github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
	gopnsensedConstant "github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	n := collector_tester.NewNetbox(t)
	k := mock_client.New()
	p := collector_tester.NewProcess(t)
	h := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	l := collector_tester.NewLease(t)

	switch peer {
	case gonetboxdConstant.Identity.Name():
		n.Failure = outage_tester.RefusedDial()
	case constant.SourceKubernetes:
		k.Failure = outage_tester.RefusedDial()
	case goprocessdConstant.Identity.Name():
		p.Failure = outage_tester.RefusedDial()
	case join.Colon(
		gooutpostdConstant.Identity.Name(),
		constant.FixtureVirtualMachineNode,
	):
		h.Failure = outage_tester.RefusedDial()
	case gopnsensedConstant.Identity.Name():
		l.Failure = outage_tester.RefusedDial()
	}

	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "atlas.db")
	o.Retention = constant.FixtureRetention
	o.Interval = constant.FixtureInterval
	o.ProcessPlace = constant.FixtureVirtualMachineNode
	o.OutpostHosts = []string{constant.FixtureVirtualMachineNode}
	o.Netbox = func() netbox.Client { return n }
	o.Kubernetes = func() kubernetes.KubernetesClient { return k }
	o.Process = func() process.Client { return p }
	o.Outpost = func(string) outpost.Client { return h }
	o.Lease = func() opnsense.Client { return l }
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
		model_context_server.WithSignOn(),
	)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
