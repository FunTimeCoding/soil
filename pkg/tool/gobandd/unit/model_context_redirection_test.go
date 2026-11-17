package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"testing"
)

func TestEnableRedirectionToolNamesWhatItSwitchedOn(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	s.Alfa().SetRedirection(false, 3)
	c := s.Tools()
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
