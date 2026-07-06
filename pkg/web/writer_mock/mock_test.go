package writer_mock

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	m := New()
	m.Header().Set("a", "b")
	assert.Any(t, http.Header{"A": {"b"}}, m.Headers)
}
