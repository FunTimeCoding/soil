package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"net/http"
	"testing"
)

func TestRestartingOpnsenseIsUnavailable(t *testing.T) {
	s := refusing(http.StatusServiceUnavailable)
	defer s.Close()
	_, e := client(t, s).Leases("")
	assert.True(t, unavailable.Is(e))
}
