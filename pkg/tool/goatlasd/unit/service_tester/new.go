package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	kubernetesCollector "github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	processCollector "github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	atlas "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	gopnsensedConstant "github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := collector_tester.NewNetbox(t)
	k := mock_client.New()
	p := collector_tester.NewProcess(t)
	d := collector_tester.NewLease(t)
	r := memory.New()
	l := logger.New(context.Background())
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "atlas.db")
	o.Retention = constant.FixtureRetention
	o.Interval = constant.FixtureInterval
	o.ProcessPlace = constant.FixtureVirtualMachineNode
	o.Netbox = func() netbox.Client { return n }
	o.Kubernetes = func() kubernetes.KubernetesClient { return k }
	o.Process = func() process.Client { return p }
	o.Lease = func() opnsense.Client { return d }
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(m)
	t.Cleanup(func() { relational.CloseMapper(m) })
	ng := guard(t, gonetboxdConstant.Identity.Name(), n.Address(), n, l, r)
	kg := guard(t, constant.SourceKubernetes, k.Configuration().Host, k, l, r)
	pg := guard(t, goprocessdConstant.Identity.Name(), p.Address(), p, l, r)
	dg := guard(t, gopnsensedConstant.Identity.Name(), d.Address(), d, l, r)

	return &Tester{
		Option: o,
		Service: service.New(
			store.New(m),
			n,
			ng,
			[]atlas.Collector{
				kubernetesCollector.New(k, kg),
				label.New(n, ng),
				processCollector.New(p, pg, o.ProcessPlace),
			},
			lease.New(d, dg),
			o.Retention,
		),
		Netbox:     n,
		Kubernetes: k,
		Lease:      d,
		Process:    p,
		Reporter:   r,
	}
}
