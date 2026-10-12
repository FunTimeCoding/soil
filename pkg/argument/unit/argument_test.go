package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	relationalConstant "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"path/filepath"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "arm", constant.Arm)
	assert.String(t, "bulk", constant.Bulk)
	assert.String(t, "cluster", constant.Cluster)
	assert.String(t, "container", constant.Container)
	assert.String(t, "context", constant.Context)
	assert.String(t, "core", constant.Core)
	assert.String(t, "dense", constant.Dense)
	assert.String(t, "disk", constant.Disk)
	assert.String(t, "download", constant.Download)
	assert.String(t, "end", constant.End)
	assert.String(t, "file", constant.File)
	assert.String(t, "group", constant.Group)
	assert.String(t, "hardware", constant.Hardware)
	assert.String(t, "information", constant.Information)
	assert.String(t, "interactive", constant.Interactive)
	assert.String(t, "investigate", constant.Investigate)
	assert.String(t, "issue", constant.Issue)
	assert.String(t, "key", constant.Key)
	assert.String(t, "log", constant.Log)
	assert.String(t, "loop", constant.Loop)
	assert.String(t, "maintainer-mail", constant.MaintainerMail)
	assert.String(t, "maintainer-name", constant.MaintainerName)
	assert.String(t, "memory", constant.Memory)
	assert.String(t, "message", constant.Message)
	assert.String(t, "migrate", constant.Migrate)
	assert.String(t, "mine", constant.Mine)
	assert.String(t, "name", constant.Name)
	assert.String(t, "node", constant.Node)
	assert.String(t, "old", constant.Old)
	assert.String(t, "on-upgrade", constant.OnUpgrade)
	assert.String(t, "package", constant.Package)
	assert.String(t, "parent", constant.Parent)
	assert.String(t, "password", constant.Password)
	assert.String(t, "pod", constant.Pod)
	assert.String(t, "pretend", constant.Pretend)
	assert.String(t, "reset", constant.Reset)
	assert.String(t, "retry", constant.Retry)
	assert.String(t, "run", constant.Run)
	assert.String(t, "runbook", constant.Runbook)
	assert.String(t, "schedule", constant.Schedule)
	assert.String(t, "source", constant.Source)
	assert.String(t, "space", constant.Space)
	assert.String(t, "statistic", constant.Statistic)
	assert.String(t, "summary", constant.Summary)
	assert.String(t, "systemd-unit", constant.SystemdUnit)
	assert.String(t, "team", constant.Team)
	assert.String(t, "title", constant.Title)
	assert.String(t, "topic", constant.Topic)
	assert.String(t, "unit", constant.Unit)
	assert.String(t, "unknown", constant.Unknown)
}

func TestLiteDefault(t *testing.T) {
	t.Setenv(relationalConstant.LitePathEnvironment, "")
	name := "gotest-lite-probe"
	a := argument.NewInstance(identity.New(name, "test tool", name))
	a.Lite()
	assert.Nil(t, a.ParseArguments(nil))
	expected := filepath.Join(
		system.StorageDirectory(name, false),
		join.Empty(name, soil.LiteExtension),
	)
	assert.String(t, expected, a.GetString(constant.Lite))
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
	assert.String(t, "/somewhere/custom.sqlite", a.GetString(constant.Lite))
}

func TestLiteFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(relationalConstant.LitePathEnvironment, "/somewhere/custom.sqlite")
	a := testInstance(t)
	a.Lite()
	assert.Nil(t, a.ParseArguments([]string{"--lite", "/explicit/flag.sqlite"}))
	assert.String(t, "/explicit/flag.sqlite", a.GetString(constant.Lite))
}

func TestDatabaseDefaults(t *testing.T) {
	t.Setenv(relationalConstant.LitePathEnvironment, "")
	t.Setenv(relationalConstant.PostgresLocatorEnvironment, "")
	a := testInstance(t)
	a.Database()
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(t, "", a.GetString(constant.Postgres))
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
		a.GetString(constant.Postgres),
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
		a.GetString(constant.Postgres),
	)
}

func TestNoPositionalsAccepts(t *testing.T) {
	a := testInstance(t)
	assert.Nil(t, a.ParseArguments(nil))
	a.NoPositionals("hint")
}

func TestNoPositionalsAcceptsFlags(t *testing.T) {
	a := testInstance(t)
	a.String(constant.File, "Procfile", constant.Path)
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
	assert.Integer(t, 8080, a.GetInteger(constant.Port))
	assert.String(t, "127.0.0.1", a.GetString(constant.BindAddress))
	assert.String(t, "127.0.0.1:8080", a.Address())
}

func TestWebEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(webConstant.PortEnvironment, "9000")
	t.Setenv(webConstant.BindEnvironment, "0.0.0.0")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(t, 9000, a.GetInteger(constant.Port))
	assert.String(t, "0.0.0.0", a.GetString(constant.BindAddress))
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
	assert.Integer(t, 7000, a.GetInteger(constant.Port))
	assert.String(t, "192.168.0.1", a.GetString(constant.BindAddress))
	assert.String(t, "192.168.0.1:7000", a.Address())
}
