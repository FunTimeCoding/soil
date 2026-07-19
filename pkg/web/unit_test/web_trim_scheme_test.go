package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTrimScheme(t *testing.T) {
	assert.String(
		t,
		"localhost",
		web.TrimScheme(locator.New(constant.Localhost).String()),
	)
	assert.String(
		t,
		"localhost",
		web.TrimScheme(locator.New(constant.Localhost).Insecure().String()),
	)
}
