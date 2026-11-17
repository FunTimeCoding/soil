package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/tool_tester"
	"strings"
	"testing"
)

func TestEmptyRequiredStringRefused(t *testing.T) {
	c := base.New(t).Tools()
	var accepted []string

	for _, tool := range c.ListTools() {
		schema := tool_tester.NewSchema(tool.InputSchema)

		for _, empty := range schema.Required {
			if schema.Properties[empty].Type != "string" {
				continue
			}

			arguments := map[string]any{}

			for _, name := range schema.Required {
				switch schema.Properties[name].Type {
				case "string":
					arguments[name] = "alfa"
				case "number", "integer":
					arguments[name] = 1
				case "array":
					arguments[name] = []any{}
				case "object":
					arguments[name] = map[string]any{}
				case "boolean":
					arguments[name] = false
				}
			}

			arguments[empty] = ""
			message, e := c.CallToolError(tool.Name, arguments)

			if e != nil || !strings.Contains(message, "required") {
				accepted = append(accepted, join.Space(tool.Name, empty))
			}
		}
	}

	assert.Strings(t, nil, accepted)
}
