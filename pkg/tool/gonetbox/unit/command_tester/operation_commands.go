package command_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"regexp"
	"strings"
)

func (o *Tester) OperationCommands() []string {
	o.t.Helper()
	s, e := server.GetSpec()
	assert.FatalOnError(o.t, e)
	word := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	var result []string

	for _, p := range s.Paths.Map() {
		for _, operation := range p.Operations() {
			name := strings.ToLower(
				word.ReplaceAllString(operation.OperationID, "$1-$2"),
			)

			if name == "get-ansible-inventory" {
				name = "ansible-inventory"
			}

			result = append(result, name)
		}
	}

	return result
}
