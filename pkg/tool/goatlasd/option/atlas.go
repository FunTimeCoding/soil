package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Atlas struct {
	option.Web
	Metric          option.Web
	LitePath        string
	PostgresLocator string
	Interval        time.Duration
	Retention       time.Duration
	ProcessPlace    string
	OutpostHosts    []string
	Netbox          func() netbox.Client
	Kubernetes      func() kubernetes.KubernetesClient
	Lease           func() opnsense.Client
	Process         func() process.Client
	Outpost         func(host string) outpost.Client
}
