package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func TestSchemePrefix(t *testing.T) {
	assert.String(t, "https://", web.SchemePrefix(true))
	// noinspection HttpUrlsUsage
	assert.String(t, "http://", web.SchemePrefix(false))
}
