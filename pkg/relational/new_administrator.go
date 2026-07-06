package relational

import (
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewAdministrator() *Database {
	return New(environment.Required(postgres.AdministratorLocatorEnvironment))
}
