package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"testing"
)

func TestAnsibleInventoryToolGroupsByPlaybookField(t *testing.T) {
	s := base.New(t)
	s.Source.SeedVirtualMachine("charlie")
	c := s.Tools()
	c.MustCallToolNotation(
		constant.SetVirtualCustomField,
		map[string]any{
			constant.VirtualMachine: "charlie",
			constant.Field:          constant.AnsibleField,
			constant.Values:         []any{"echo"},
		},
	)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                charlie: {}
`,
		c.MustCallTool(constant.AnsibleInventory, map[string]any{}),
	)
}

func TestConfigurationContextToolWeightDefault(t *testing.T) {
	s := base.New(t)
	var result client.ConfigurationContext
	notation.MustDecode(
		s.Tools().MustCallTool(
			constant.CreateConfigurationContext,
			map[string]any{
				"name":           "alfa",
				constant.Payload: map[string]any{"bravo": "charlie"},
			},
		),
		&result,
		false,
	)
	assert.Integer(t, 1000, int(result.Weight))
}
