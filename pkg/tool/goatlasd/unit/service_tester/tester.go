package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	opnsense "github.com/funtimecoding/soil/pkg/tool/gopnsensed/client/mock_client"
	process "github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
)

type Tester struct {
	Service    *service.Service
	Netbox     *netbox.Client
	Kubernetes *kubernetes.Client
	Lease      *opnsense.Client
	Process    *process.Client
	Reporter   *memory.Memory
}
