package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/spy_writer"
	"github.com/funtimecoding/soil/pkg/web/writer_mock"
	"net/http"
	"testing"
)

func TestWriteOkay(t *testing.T) {
	w := spy_writer.New()
	assert.Integer(t, 4, web.WriteOkay(w, upper.Alfa))
	assert.Any(t, []byte("Alfa"), w.Written)
	assert.Integer(t, 200, w.StatusCode)
}

func TestTextHeader(t *testing.T) {
	m := writer_mock.New()
	web.TextHeader(m)
	assert.Any(
		t,
		http.Header{"Content-Type": {"text/plain"}},
		m.Headers,
	)
}

func TestObjectHeader(t *testing.T) {
	m := writer_mock.New()
	web.ObjectHeader(m)
	assert.Any(
		t,
		http.Header{"Content-Type": {"application/json"}},
		m.Headers,
	)
}

func TestMock(t *testing.T) {
	m := writer_mock.New()
	m.Header().Set("a", "b")
	assert.Any(t, http.Header{"A": {"b"}}, m.Headers)
}
