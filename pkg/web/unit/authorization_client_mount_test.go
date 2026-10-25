package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestAMountedClientSendsABareSessionToSignIn(t *testing.T) {
	w := served(newMountedMux(t), "/target")
	assert.Integer(t, http.StatusFound, w.Code)
	assert.String(t, "/sign-in", w.Header().Get("Location"))
}

func TestAMountedClientAnswersTheCallback(t *testing.T) {
	w := served(newMountedMux(t), constant.CallbackPath)
	assert.Integer(t, http.StatusBadRequest, w.Code)
	assert.StringContains(t, "missing flow state", w.Body.String())
}

func TestAMountedClientSignsOut(t *testing.T) {
	w := served(newMountedMux(t), constant.SignOutPath)
	assert.Integer(t, http.StatusFound, w.Code)
	assert.String(t, "/", w.Header().Get("Location"))
	assert.StringContains(t, "Max-Age=0", w.Header().Get("Set-Cookie"))
}
