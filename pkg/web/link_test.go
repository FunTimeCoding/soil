package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"https://localhost:9000",
		Link(constant.Localhost, 9000, true),
	)
	assert.String(
		t,
		"https://localhost",
		Link(constant.Localhost, 0, true),
	)
	assert.String(
		t,
		"http://localhost:9000",
		Link(constant.Localhost, 9000, false),
	)
	assert.String(
		t,
		"http://localhost",
		Link(constant.Localhost, 0, false),
	)
}
