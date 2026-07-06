package relational

import (
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Database {
	return New(environment.Required(postgres.LocatorEnvironment))
}
