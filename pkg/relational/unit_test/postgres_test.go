package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "POSTGRES_LOCATOR", postgres.LocatorEnvironment)
	assert.String(t, "psql", postgres.Command)
	assert.String(t, "--username", postgres.UserArgument)
	assert.String(t, "--command", postgres.CommandArgument)
	assert.String(t, "--file", postgres.FileArgument)
	assert.String(t, "--echo-all", postgres.EchoAllFlag)
	assert.String(t, "pg_dump", postgres.DumpCommand)
	assert.String(t, "postgres", postgres.DialectName)
}
