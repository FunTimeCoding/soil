package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSchemePrefix(t *testing.T) {
	assert.String(t, "https://", SchemePrefix(true))
	// noinspection HttpUrlsUsage
	assert.String(t, "http://", SchemePrefix(false))
}
