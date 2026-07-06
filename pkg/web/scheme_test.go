package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestScheme(t *testing.T) {
	assert.String(t, "https", Scheme(true))
	assert.String(t, "http", Scheme(false))
}
