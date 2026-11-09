package dump

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"strconv"
)

func connection(
	i *instance.Instance,
	database string,
) []string {
	return []string{
		"--host",
		i.Host,
		"--port",
		strconv.Itoa(i.Port),
		"--username",
		i.User,
		"--dbname",
		database,
		"--no-password",
	}
}
