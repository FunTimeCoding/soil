package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/terminal/terminal_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestCommandCoversEveryOperation(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	var commands []string

	for _, o := range gonetbox.New(c, terminal_tester.New().Terminal).Commands() {
		commands = append(commands, o.Name())
	}

	s, e := server.GetSpec()
	assert.FatalOnError(t, e)
	word := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	var missing []string

	for _, p := range s.Paths.Map() {
		for _, o := range p.Operations() {
			name := strings.ToLower(
				word.ReplaceAllString(o.OperationID, "$1-$2"),
			)

			if name == "get-ansible-inventory" {
				name = "ansible-inventory"
			}

			if !slices.Contains(commands, name) {
				missing = append(missing, name)
			}
		}
	}

	slices.Sort(missing)
	assert.Strings(t, nil, missing)
}
