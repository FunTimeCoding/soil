package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
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
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "atlas.db")
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://atlas.example.org"
	o.Retention = constant.FixtureRetention
	o.ProcessPlace = constant.FixtureVirtualMachineNode
	o.Netbox = func() netbox.Client { return n }
	o.Kubernetes = func() kubernetes.Client { return k }
	o.Process = func() process.Client { return p }
	o.Lease = func() opnsense.Client { return d }
	result := &Tester{
		Service:    service.New(o, logger.New(context.Background()), r),
		Netbox:     n,
		Kubernetes: k,
		Lease:      d,
		Process:    p,
		Reporter:   r,
	}
	t.Cleanup(result.Service.Close)

	return result
}
