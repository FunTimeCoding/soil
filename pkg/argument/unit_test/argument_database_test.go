package unit_test

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	relational "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestLiteDefault(t *testing.T) {
	t.Setenv(relational.LitePathEnvironment, "")
	name := "gotest-lite-probe"
	a := argument.NewInstance(identity.New(name, "test tool", name))
	a.Lite()
	assert.Nil(t, a.ParseArguments(nil))
	expected := filepath.Join(
		system.StorageDirectory(name, false),
		join.Empty(name, libraryConstant.LiteExtension),
	)
	assert.String(t, expected, a.GetString(argumentConstant.Lite))
	assert.False(
		t,
		system.DirectoryExists(system.StorageDirectory(name, false)),
	)
}

func TestLiteEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(relational.LitePathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(
		t,
		"/somewhere/custom.sqlite",
		a.GetString(argumentConstant.Lite),
	)
}

func TestLiteFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(relational.LitePathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(
		t,
		a.ParseArguments([]string{"--lite", "/explicit/flag.sqlite"}),
	)
	assert.String(
		t,
		"/explicit/flag.sqlite",
		a.GetString(argumentConstant.Lite),
	)
}

func TestDatabaseDefaults(t *testing.T) {
	t.Setenv(relational.LitePathEnvironment, "")
	t.Setenv(relational.PostgresLocatorEnvironment, "")
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "", a.GetString(argumentConstant.Postgres))
}

func TestDatabaseEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(
		relational.PostgresLocatorEnvironment,
		"postgres://env@localhost/env",
	)
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(
		t,
		"postgres://env@localhost/env",
		a.GetString(argumentConstant.Postgres),
	)
}

func TestDatabaseFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(
		relational.PostgresLocatorEnvironment,
		"postgres://env@localhost/env",
	)
	a := testInstance(t)
	a.Database()
	assert.Nil(
		t,
		a.ParseArguments(
			[]string{"--postgres", "postgres://flag@localhost/flag"},
		),
	)
	assert.String(
		t,
		"postgres://flag@localhost/flag",
		a.GetString(argumentConstant.Postgres),
	)
}
