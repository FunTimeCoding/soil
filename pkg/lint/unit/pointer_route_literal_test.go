package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestRouteLiteral(t *testing.T) {
	assert.String(t, "/widgets", pointer.RouteLiteral("/widgets"))
	assert.String(
		t,
		"/memories/",
		pointer.RouteLiteral("/memories/{identifier}"),
	)
	assert.String(t, "/api/", pointer.RouteLiteral("/api/..."))
	assert.String(t, "/", pointer.RouteLiteral("/"))
}

func TestContainsLiteral(t *testing.T) {
	assert.True(t, pointer.ContainsLiteral(`Path = "/widgets"`, "/widgets"))
	assert.True(
		t,
		pointer.ContainsLiteral(
			"path: /metrics/cadvisor\n",
			"/metrics/cadvisor",
		),
	)
	assert.True(
		t,
		pointer.ContainsLiteral(`Handle("/memories/", h)`, "/memories/"),
	)
	assert.True(
		t,
		pointer.ContainsLiteral(
			"image: ghcr.io/example/alpha:v1",
			"ghcr.io/example/alpha",
		),
	)
	assert.False(t, pointer.ContainsLiteral(`"/widgets-detail"`, "/widgets"))
	assert.False(t, pointer.ContainsLiteral(`"/api/widgets"`, "/widgets"))
	assert.False(
		t,
		pointer.ContainsLiteral(
			"image: ghcr.io/example/alphabet:v1",
			"ghcr.io/example/alpha",
		),
	)
}
