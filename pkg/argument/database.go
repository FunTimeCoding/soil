package argument

import (
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Database() {
	i.Lite()
	i.String(
		Postgres,
		environment.Optional(postgres.LocatorEnvironment),
		postgres.LocatorUsage,
	)
}
