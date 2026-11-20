package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"slices"
	"testing"
)

func TestCommandCoversEveryOperation(t *testing.T) {
	commands := command_tester.New(t).CommandNames()
	var missing []string

	for _, name := range command_tester.New(t).OperationCommands() {
		if !slices.Contains(commands, name) {
			missing = append(missing, name)
		}
	}

	slices.Sort(missing)
	assert.Strings(t, nil, missing)
}

func TestEveryCommandHasOperation(t *testing.T) {
	operations := command_tester.New(t).OperationCommands()
	var extra []string

	for _, name := range command_tester.New(t).CommandNames() {
		if !slices.Contains(operations, name) {
			extra = append(extra, name)
		}
	}

	slices.Sort(extra)
	assert.Strings(t, nil, extra)
}
