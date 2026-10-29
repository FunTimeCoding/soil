package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/closing_source"
	"testing"
)

func TestServiceClosesTheClientItBuilt(t *testing.T) {
	c := closing_source.New()
	o := option.New()
	o.Source = func() face.OpnsenseSource { return c }
	service.New(o).Close()
	assert.Integer(t, 1, c.Closed())
}
