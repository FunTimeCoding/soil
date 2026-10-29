package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"net/http"
	"testing"
)

func TestForbiddenClassifiesByType(t *testing.T) {
	e := forbidden.Format("not armed: %s", "close")
	assert.String(t, "not armed: close", e.Error())
	assert.True(t, forbidden.Is(e))
	assert.False(t, forbidden.Is(errors.New("not armed: close")))
	assert.False(t, forbidden.Is(nil))
}

func TestForbiddenAnswersAnUncaptured403(t *testing.T) {
	e := fmt.Errorf("close: %w", forbidden.Format("not armed"))
	assert.String(t, "forbidden", classify.Class(e))
	assert.Integer(t, http.StatusForbidden, classify.Answer(e).Status)
	assert.False(t, classify.Reportable(e))
}
