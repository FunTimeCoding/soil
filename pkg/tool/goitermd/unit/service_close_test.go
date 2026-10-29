package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/closing_client"
	"testing"
)

func TestServiceClosesTheClientItBuilt(t *testing.T) {
	c := closing_client.New()
	o := option.New()
	o.Source = func() face.ItermSource { return c }
	service.New(o).Close()
	assert.Integer(t, 1, c.Closed())
}
