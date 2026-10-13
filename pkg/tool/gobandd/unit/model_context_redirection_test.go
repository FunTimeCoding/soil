package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"testing"
)

func TestEnableRedirectionToolNamesWhatItSwitchedOn(t *testing.T) {
	s := base.New(t)
	s.Alfa.SetRedirection(false, 3)
	c := s.Connect(t)
	defer c.Close()
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)
	assert.String(
		t,
		"enabled listener, screen service on alfa",
		c.MustCallTool(constant.EnableRedirection, nil),
	)
	assert.String(
		t,
		"redirection already enabled on alfa",
		c.MustCallTool(constant.EnableRedirection, nil),
	)
}
