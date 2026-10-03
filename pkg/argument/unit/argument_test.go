package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	relationalConstant "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"path/filepath"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "arm", argumentConstant.Arm)
	assert.String(t, "bulk", argumentConstant.Bulk)
	assert.String(t, "cluster", argumentConstant.Cluster)
	assert.String(t, "container", argumentConstant.Container)
	assert.String(t, "context", argumentConstant.Context)
	assert.String(t, "core", argumentConstant.Core)
	assert.String(t, "dense", argumentConstant.Dense)
	assert.String(t, "disk", argumentConstant.Disk)
	assert.String(t, "download", argumentConstant.Download)
	assert.String(t, "end", argumentConstant.End)
	assert.String(t, "file", argumentConstant.File)
	assert.String(t, "group", argumentConstant.Group)
	assert.String(t, "hardware", argumentConstant.Hardware)
	assert.String(t, "information", argumentConstant.Information)
	assert.String(t, "interactive", argumentConstant.Interactive)
	assert.String(t, "investigate", argumentConstant.Investigate)
	assert.String(t, "issue", argumentConstant.Issue)
	assert.String(t, "key", argumentConstant.Key)
	assert.String(t, "log", argumentConstant.Log)
	assert.String(t, "loop", argumentConstant.Loop)
	assert.String(t, "maintainer-mail", argumentConstant.MaintainerMail)
	assert.String(t, "maintainer-name", argumentConstant.MaintainerName)
	assert.String(t, "memory", argumentConstant.Memory)
	assert.String(t, "message", argumentConstant.Message)
	assert.String(t, "migrate", argumentConstant.Migrate)
	assert.String(t, "mine", argumentConstant.Mine)
	assert.String(t, "name", argumentConstant.Name)
	assert.String(t, "node", argumentConstant.Node)
	assert.String(t, "old", argumentConstant.Old)
	assert.String(t, "on-upgrade", argumentConstant.OnUpgrade)
	assert.String(t, "package", argumentConstant.Package)
	assert.String(t, "parent", argumentConstant.Parent)
	assert.String(t, "password", argumentConstant.Password)
	assert.String(t, "pod", argumentConstant.Pod)
	assert.String(t, "pretend", argumentConstant.Pretend)
	assert.String(t, "reset", argumentConstant.Reset)
	assert.String(t, "retry", argumentConstant.Retry)
	assert.String(t, "run", argumentConstant.Run)
	assert.String(t, "runbook", argumentConstant.Runbook)
	assert.String(t, "schedule", argumentConstant.Schedule)
	assert.String(t, "source", argumentConstant.Source)
	assert.String(t, "space", argumentConstant.Space)
	assert.String(t, "statistic", argumentConstant.Statistic)
	assert.String(t, "summary", argumentConstant.Summary)
	assert.String(t, "systemd-unit", argumentConstant.SystemdUnit)
	assert.String(t, "team", argumentConstant.Team)
	assert.String(t, "title", argumentConstant.Title)
	assert.String(t, "topic", argumentConstant.Topic)
	assert.String(t, "unit", argumentConstant.Unit)
	assert.String(t, "unknown", argumentConstant.Unknown)
}

func TestLiteDefault(t *testing.T) {
	t.Setenv(relationalConstant.LitePathEnvironment, "")
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
	t.Setenv(relationalConstant.LitePathEnvironment, "/somewhere/custom.sqlite")
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
	t.Setenv(relationalConstant.LitePathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(t, a.ParseArguments([]string{"--lite", "/explicit/flag.sqlite"}))
	assert.String(
		t,
		"/explicit/flag.sqlite",
		a.GetString(argumentConstant.Lite),
	)
}

func TestDatabaseDefaults(t *testing.T) {
	t.Setenv(relationalConstant.LitePathEnvironment, "")
	t.Setenv(relationalConstant.PostgresLocatorEnvironment, "")
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "", a.GetString(argumentConstant.Postgres))
}

func TestDatabaseEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(
		relationalConstant.PostgresLocatorEnvironment,
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
		relationalConstant.PostgresLocatorEnvironment,
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

func TestNoPositionalsAccepts(t *testing.T) {
	a := testInstance(t)
	assert.Nil(t, a.ParseArguments(nil))
	a.NoPositionals("hint")
}

func TestNoPositionalsAcceptsFlags(t *testing.T) {
	a := testInstance(t)
	a.String(argumentConstant.File, "Procfile", argumentConstant.Path)
	assert.Nil(t, a.ParseArguments([]string{"--file", "Other"}))
	a.NoPositionals("hint")
}

func TestPositionalOutOfBoundsIsEmpty(t *testing.T) {
	assert.String(t, "", argument.Positional(99))
}

func TestWebDefaults(t *testing.T) {
	t.Setenv(webConstant.PortEnvironment, "")
	t.Setenv(webConstant.BindEnvironment, "")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(t, 8080, a.GetInteger(argumentConstant.Port))
	assert.String(t, "127.0.0.1", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "127.0.0.1:8080", a.Address())
}

func TestWebEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(webConstant.PortEnvironment, "9000")
	t.Setenv(webConstant.BindEnvironment, "0.0.0.0")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(t, 9000, a.GetInteger(argumentConstant.Port))
	assert.String(t, "0.0.0.0", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "0.0.0.0:9000", a.Address())
}

func TestWebFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(webConstant.PortEnvironment, "9000")
	t.Setenv(webConstant.BindEnvironment, "0.0.0.0")
	a := testInstance(t)
	a.Web()
	assert.Nil(
		t,
		a.ParseArguments(
			[]string{"--port", "7000", "--bind-address", "192.168.0.1"},
		),
	)
	assert.Integer(t, 7000, a.GetInteger(argumentConstant.Port))
	assert.String(t, "192.168.0.1", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "192.168.0.1:7000", a.Address())
}
