package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"https://localhost:9000",
		web.Link(constant.Localhost, 9000, true),
	)
	assert.String(
		t,
		"https://localhost",
		web.Link(constant.Localhost, 0, true),
	)
	assert.String(
		t,
		"http://localhost:9000",
		web.Link(constant.Localhost, 9000, false),
	)
	assert.String(
		t,
		"http://localhost",
		web.Link(constant.Localhost, 0, false),
	)
}
