package postgres

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "POSTGRES_LOCATOR", LocatorEnvironment)
	assert.String(t, "psql", Command)
	assert.String(t, "--username", UserArgument)
	assert.String(t, "--command", CommandArgument)
	assert.String(t, "--file", FileArgument)
	assert.String(t, "--echo-all", EchoAllFlag)
	assert.String(t, "pg_dump", DumpCommand)
	assert.String(t, "postgres", DialectName)
}
