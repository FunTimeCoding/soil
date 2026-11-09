package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"strings"
	"testing"
)

func TestComponentStyleComesBeforeDaemonsOwn(t *testing.T) {
	page := web_tester.RenderBrand(
		layout.New(identity.New("test", "test tool", "test"), "", nil).
			WithStyle(".alfa { color: red; }"),
	)
	component := strings.Index(page, constant.ComponentStyle)
	own := strings.Index(page, ".alfa { color: red; }")
	assert.True(t, component >= 0)
	assert.True(t, component < own)
	assert.Integer(t, 1, strings.Count(page, constant.ComponentStyle))
}
