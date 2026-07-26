package relational

import (
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Database {
	return New(environment.Required(constant.PostgresLocatorEnvironment))
}
