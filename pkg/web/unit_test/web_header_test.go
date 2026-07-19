package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/writer_mock"
	"net/http"
	"testing"
)

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
