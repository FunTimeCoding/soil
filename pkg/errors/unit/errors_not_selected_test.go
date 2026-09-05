package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"testing"
)

func TestNotSelectedClassifiesByType(t *testing.T) {
	e := not_selected.Format("no instance selected - %d configured", 3)
	assert.String(t, "no instance selected - 3 configured", e.Error())
	assert.True(t, not_selected.Is(e))
	assert.False(t, not_selected.Is(errors.New("no instance selected")))
	assert.False(t, not_selected.Is(nil))
}
