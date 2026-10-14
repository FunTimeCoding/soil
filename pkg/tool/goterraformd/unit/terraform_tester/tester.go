package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"testing"
)

type Tester struct {
	t        *testing.T
	Runner   *runner.Runner
	Cluster  *mock_client.Client
	Reporter *memory.Memory
	store    *store.Store
	log      string
}
