package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_tester"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

func TestLoginSavesDeviceToken(t *testing.T) {
	s := graph_server.New(t).
		Route(
			graph_tester.DevicePath(),
			http.StatusOK,
			graph_tester.DeviceCode(900),
		).
		Route(graph_tester.TokenPath(), http.StatusOK, graph_tester.Refreshed())
	store := mock_token_store.New(t, nil)
	var shown string
	assert.FatalOnError(
		t,
		graph_tester.New(t).Client(s, store).Login(
			t.Context(),
			func(a *oauth2.DeviceAuthResponse) { shown = a.UserCode },
		),
	)
	assert.String(t, "USER-1", shown)
	poll := s.Exchanges(graph_tester.TokenPath())
	assert.Count(t, 1, poll)
	assert.String(t, "device-1", poll[0].Form.Get("device_code"))
	saved := store.Saved()
	assert.Count(t, 1, saved)
	assert.String(t, "fresh-access", saved[0].AccessToken)
}

func TestDeclinedLoginInvalid(t *testing.T) {
	store, e := graph_tester.New(t).LoginRefusal(constant.DeclinedAuthorization)
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"graph sign-in not finished: authorization_declined AADSTS70000",
		e.Error(),
	)
	assert.Count(t, 0, store.Saved())
}

func TestDeniedLoginInvalid(t *testing.T) {
	store, e := graph_tester.New(t).LoginRefusal(constant.AccessDenied)
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"graph sign-in not finished: access_denied AADSTS70000",
		e.Error(),
	)
	assert.Count(t, 0, store.Saved())
}

func TestExpiredLoginCodeInvalid(t *testing.T) {
	store, e := graph_tester.New(t).LoginRefusal(constant.ExpiredDeviceCode)
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"graph sign-in not finished: expired_token AADSTS70000",
		e.Error(),
	)
	assert.Count(t, 0, store.Saved())
}

func TestExpiredDeviceCodeInvalid(t *testing.T) {
	s := graph_server.New(t).
		Route(
			graph_tester.DevicePath(),
			http.StatusOK,
			graph_tester.DeviceCode(1),
		).
		Route(
			graph_tester.TokenPath(),
			http.StatusBadRequest,
			graph_tester.Refusal("authorization_pending", "AADSTS70016"),
		)
	e := graph_tester.New(t).Client(s, mock_token_store.New(t, nil)).Login(
		t.Context(),
		func(*oauth2.DeviceAuthResponse) {},
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "graph sign-in not finished: expired_token", e.Error())
}
