package tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
)

func (t *Tester) Send(
	command string,
	arguments ...string,
) string {
	result, e := client.Send(t.SocketPath, command, arguments)
	errors.PanicOnError(e)

	return result
}
