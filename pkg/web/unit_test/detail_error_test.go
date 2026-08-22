package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestDetailErrorCarriesUpstreamContext(t *testing.T) {
	e := detail_error.New("zone already exists", "400 Bad Request")
	e.Body = []byte(`{"status":"error"}`)
	assert.String(t, "zone already exists", e.Error())
	var provider face.ContextProvider = e
	key, context := provider.ErrorContext()
	assert.String(t, "upstream", key)
	assert.String(t, "400 Bad Request", context[constant.Status].(string))
	assert.String(t, "zone already exists", context[constant.Detail].(string))
	assert.String(t, `{"status":"error"}`, context[constant.Body].(string))
}
