package command_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"regexp"
	"strings"
	"testing"
)

func OperationCommands(t *testing.T) []string {
	t.Helper()
	s, e := server.GetSpec()
	assert.FatalOnError(t, e)
	word := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	var result []string

	for _, p := range s.Paths.Map() {
		for _, o := range p.Operations() {
			name := strings.ToLower(
				word.ReplaceAllString(o.OperationID, "$1-$2"),
			)

			if name == "get-ansible-inventory" {
				name = "ansible-inventory"
			}

			result = append(result, name)
		}
	}

	return result
}
