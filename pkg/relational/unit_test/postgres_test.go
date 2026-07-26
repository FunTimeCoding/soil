package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "POSTGRES_LOCATOR", constant.PostgresLocatorEnvironment)
	assert.String(t, "psql", constant.PostgresCommand)
	assert.String(t, "--username", constant.PostgresUserArgument)
	assert.String(t, "--command", constant.PostgresCommandArgument)
	assert.String(t, "--file", constant.PostgresFileArgument)
	assert.String(t, "--echo-all", constant.PostgresEchoAllFlag)
	assert.String(t, "pg_dump", constant.PostgresDumpCommand)
	assert.String(t, "postgres", constant.PostgresDialectName)
}
