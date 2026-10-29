package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"testing"
)

type Tester struct {
	t        *testing.T
	Service  face.Provision
	Cluster  *mock_client.Client
	Reporter *memory.Memory
	log      string
}
