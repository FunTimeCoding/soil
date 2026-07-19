package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func TestScheme(t *testing.T) {
	assert.String(t, "https", web.Scheme(true))
	assert.String(t, "http", web.Scheme(false))
}
