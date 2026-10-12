package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	relationalConstant "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Database() {
	i.Lite()
	i.String(
		constant.Postgres,
		environment.Optional(relationalConstant.PostgresLocatorEnvironment),
		relationalConstant.PostgresLocatorUsage,
	)
}
