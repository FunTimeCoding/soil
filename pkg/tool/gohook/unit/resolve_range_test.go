package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"testing"
)

func TestResolveRangeExplicitWinsOverInput(t *testing.T) {
	o := option.New()
	o.Hook = "pre-push"
	o.Base = "a"
	o.Head = "b"
	o.Input = "refs/heads/main x refs/heads/main y\n"
	assert.String(t, "a..b", gohook.ResolveRange("", o).String())
}

func TestResolveRangeInputOnlyForPrePush(t *testing.T) {
	o := option.New()
	o.Hook = "post-merge"
	o.Input = "refs/heads/main x refs/heads/main y\n"
	assert.True(t, gohook.ResolveRange("", o).All)
}
