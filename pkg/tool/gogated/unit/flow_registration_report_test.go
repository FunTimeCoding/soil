package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"net/http"
	"testing"
)

func TestARefusedRegistrationIsNotReported(t *testing.T) {
	o := tester.New(t)
	status := o.RegisterRaw(t, map[string]any{"client_name": "no-redirect"})
	assert.Integer(t, http.StatusBadRequest, status)
	assert.Count(t, 0, o.Reporter().Events())
}
