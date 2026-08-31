package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/response"
	"testing"
)

func TestResponse(t *testing.T) {
	r := response.New("body", 200)
	assert.String(t, "body", r.Body)
	assert.Integer(t, 200, r.Status)
}

func TestPrintWrapperAnchors(t *testing.T) {
	console.Line()
	console.Formatted("")
	assert.NotNil(t, console.Emit)
}
