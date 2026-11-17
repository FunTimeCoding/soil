package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/flow_tester"
	"net/http"
	"testing"
)

func TestRefusedRegistrationIsNotReported(t *testing.T) {
	o := flow_tester.New(t)
	status := o.RegisterRaw(map[string]any{"client_name": "no-redirect"})
	assert.Integer(t, http.StatusBadRequest, status)
	assert.Count(t, 0, o.Reporter().Events())
}
