package postgres_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"strconv"
)

func (o *Tester) Psql(
	container string,
	i *instance.Instance,
	database string,
	sql string,
) string {
	o.t.Helper()
	r := run.New().NoPanic()
	result := r.Start(
		"podman",
		"exec",
		"--env",
		fmt.Sprintf("%s=%s", constant.PasswordEnvironment, i.Password),
		container,
		"psql",
		"--host",
		i.Host,
		"--port",
		strconv.Itoa(i.Port),
		"--username",
		i.User,
		"--dbname",
		database,
		"--tuples-only",
		"--no-align",
		"--command",
		sql,
	)
	assert.FatalOnError(o.t, r.Error)

	return result
}
