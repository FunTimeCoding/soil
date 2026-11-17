package command_tester

import (
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func CommandNames(t *testing.T) []string {
	t.Helper()
	c := mock_client.New(t)
	defer c.Close()
	var result []string

	for _, o := range gonetbox.New(c, terminal_tester.New().Terminal).Commands() {
		result = append(result, o.Name())
	}

	return result
}
