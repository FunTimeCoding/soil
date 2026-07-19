package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	assert.Any(t, &http.Client{}, web.Client())
}
