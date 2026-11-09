package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"testing"
)

func TestBackupDirectoryComesFromEnvironment(t *testing.T) {
	t.Setenv(constant.BackupDirectoryEnvironment, "/environment/backup")
	a := argument.NewInstance(constant.Identity)
	gopostgresd.Arguments(a)
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(
		t,
		"/environment/backup",
		a.GetString(constant.BackupDirectoryArgument),
	)
}

func TestBackupDirectoryFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(constant.BackupDirectoryEnvironment, "/environment/backup")
	a := argument.NewInstance(constant.Identity)
	gopostgresd.Arguments(a)
	assert.Nil(
		t,
		a.ParseArguments([]string{"--backup-directory", "/flag/backup"}),
	)
	assert.String(
		t,
		"/flag/backup",
		a.GetString(constant.BackupDirectoryArgument),
	)
}
