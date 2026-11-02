package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/service"
	"testing"
)

type Tester struct {
	t        *testing.T
	Option   *option.Terraform
	Service  *service.Service
	Cluster  *mock_client.Client
	Reporter *memory.Memory
	served   *model_context_server.Server
	log      string
	litePath string
}
