package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	assert.Any(t, &http.Client{}, Client())
}
