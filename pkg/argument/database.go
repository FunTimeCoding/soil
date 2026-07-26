package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Database() {
	i.Lite()
	i.String(
		constant.Postgres,
		environment.Optional(postgres.LocatorEnvironment),
		postgres.LocatorUsage,
	)
}
