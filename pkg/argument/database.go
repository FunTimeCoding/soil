package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	relational "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Database() {
	i.Lite()
	i.String(
		constant.Postgres,
		environment.Optional(relational.PostgresLocatorEnvironment),
		relational.PostgresLocatorUsage,
	)
}
