package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestIsDeadTag(t *testing.T) {
	assert.True(
		t,
		go_mod.IsDeadTag(
			"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
		),
	)
}

func TestIsDeadTagNegative(t *testing.T) {
	assert.False(t, go_mod.IsDeadTag("go: module not found"))
}

func TestParseDeadTag(t *testing.T) {
	mod, version := go_mod.ParseDeadTag(
		"go: github.com/funtimecoding/soil@v0.10.307: reading github.com/funtimecoding/soil/go.mod at revision v0.10.307: unknown revision v0.10.307",
	)
	assert.String(t, "github.com/funtimecoding/soil", mod)
	assert.String(t, "v0.10.307", version)
}

func TestParseDeadTagNoMatch(t *testing.T) {
	mod, version := go_mod.ParseDeadTag("go: module not found")
	assert.String(t, "", mod)
	assert.String(t, "", version)
}

func TestReadProjectFromRelativePath(t *testing.T) {
	t.Chdir(t.TempDir())
	system.SaveFile(constant.ModFile, "module a\n\ngo 1.25.0\n")
	assert.String(
		t,
		"1.25.0",
		go_mod.ReadProject(soil.CurrentDirectory, "1.27.1").Version,
	)
}
