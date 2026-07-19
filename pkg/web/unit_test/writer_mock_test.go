package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/writer_mock"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	m := writer_mock.New()
	m.Header().Set("a", "b")
	assert.Any(t, http.Header{"A": {"b"}}, m.Headers)
}
