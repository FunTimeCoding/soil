package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTrimScheme(t *testing.T) {
	assert.String(
		t,
		"localhost",
		TrimScheme(locator.New(constant.Localhost).String()),
	)
	assert.String(
		t,
		"localhost",
		TrimScheme(locator.New(constant.Localhost).Insecure().String()),
	)
}
