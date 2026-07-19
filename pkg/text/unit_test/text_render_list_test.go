package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text"
	"testing"
)

func TestRenderList(t *testing.T) {
	assert.String(
		t,
		text.RenderList("title", []string{"element"}),
		"title:\nelement",
	)
}
