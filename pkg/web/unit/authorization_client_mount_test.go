package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/unit/issuer_server"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"net/http"
	"testing"
)

func TestMountedClientSendsBareSessionToSignIn(t *testing.T) {
	w := web_tester.Served(issuer_server.New(t).Mux(), "/target")
	assert.Integer(t, http.StatusFound, w.Code)
	assert.String(t, "/sign-in", w.Header().Get("Location"))
}

func TestMountedClientAnswersCallback(t *testing.T) {
	w := web_tester.Served(issuer_server.New(t).Mux(), constant.CallbackPath)
	assert.Integer(t, http.StatusBadRequest, w.Code)
	assert.StringContains(t, "missing flow state", w.Body.String())
}

func TestMountedClientSignsOut(t *testing.T) {
	w := web_tester.Served(issuer_server.New(t).Mux(), constant.SignOutPath)
	assert.Integer(t, http.StatusFound, w.Code)
	assert.String(t, "/", w.Header().Get("Location"))
	assert.StringContains(t, "Max-Age=0", w.Header().Get("Set-Cookie"))
}
