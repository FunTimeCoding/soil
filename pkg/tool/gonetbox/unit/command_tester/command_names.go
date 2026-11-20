package command_tester

import (
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
)

func (o *Tester) CommandNames() []string {
	o.t.Helper()
	c := mock_client.New(o.t)
	defer c.Close()
	var result []string

	for _, command := range gonetbox.New(
		c,
		terminal_tester.New().Terminal,
	).Commands() {
		result = append(result, command.Name())
	}

	return result
}
