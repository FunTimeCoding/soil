package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/service"
	"testing"
)

type Tester struct {
	t          *testing.T
	Option     *option.Terraform
	Cluster    *mock_client.Client
	Reporter   *memory.Memory
	service    *service.Service
	hostname   string
	unlockExit int
	plan       string
	downstream []*provisionMock.Client
	log        string
	litePath   string
}
