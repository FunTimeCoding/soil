package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
	"time"
)

type Atlas struct {
	Address          string
	MetricAddress    string
	ServiceTokens    []string
	LitePath         string
	PostgresLocator  string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	Interval         time.Duration
	Retention        time.Duration
	ProcessPlace     string
	OutpostHosts     []string
	Netbox           func() netbox.Client
	Kubernetes       func() kubernetes.Client
	Lease            func() opnsense.Client
	Process          func() process.Client
	Outpost          func(host string) outpost.Client
}
