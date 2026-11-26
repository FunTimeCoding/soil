package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/option"
	"testing"
)

func TestOptionDefaults(t *testing.T) {
	o := option.New()
	assert.String(t, "goladybirdd", o.Profile)
	assert.True(t, o.Headless)
	assert.NotNil(t, o.Client)
}
