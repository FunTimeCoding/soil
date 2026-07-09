package argument

import (
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/relational/postgres"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestLiteDefault(t *testing.T) {
	t.Setenv(constant.PathEnvironment, "")
	name := "gotest-lite-probe"
	a := NewInstance(identity.New(name, "test tool", name))
	a.Lite()
	assert.Nil(t, a.ParseArguments(nil))
	expected := filepath.Join(
		system.StorageDirectory(name, false),
		join.Empty(name, libraryConstant.LiteExtension),
	)
	assert.String(t, expected, a.GetString(Lite))
	assert.False(t, system.DirectoryExists(system.StorageDirectory(name, false)))
}

func TestLiteEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(constant.PathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "/somewhere/custom.sqlite", a.GetString(Lite))
}

func TestLiteFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(constant.PathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(t, a.ParseArguments([]string{"--lite", "/explicit/flag.sqlite"}))
	assert.String(t, "/explicit/flag.sqlite", a.GetString(Lite))
}

func TestDatabaseDefaults(t *testing.T) {
	t.Setenv(constant.PathEnvironment, "")
	t.Setenv(postgres.LocatorEnvironment, "")
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "", a.GetString(Postgres))
}

func TestDatabaseEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(postgres.LocatorEnvironment, "postgres://env@localhost/env")
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "postgres://env@localhost/env", a.GetString(Postgres))
}

func TestDatabaseFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(postgres.LocatorEnvironment, "postgres://env@localhost/env")
	a := testInstance(t)
	a.Database()
	assert.Nil(
		t,
		a.ParseArguments(
		[]string{"--postgres", "postgres://flag@localhost/flag"},
	))
	assert.String(t, "postgres://flag@localhost/flag", a.GetString(Postgres))
}
