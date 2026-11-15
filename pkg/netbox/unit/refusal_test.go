package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestRejectedWriteIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"slug":["site with this slug already exists."],"name":["required"]}`,
		),
	)
	_, e := netbox_tester.New(t, s.URL).CreateSite("alfa")
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"name: required; slug: site with this slug already exists.",
		e.Error(),
	)
}

func TestMissingObjectIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"detail":"No Site matches the given query."}`,
		),
	)
	_, e := netbox_tester.New(t, s.URL).Sites()
	assert.True(t, not_found.Is(e))
	assert.String(t, "No Site matches the given query.", e.Error())
}

func TestConflictIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(
			`{"detail":"Unable to delete: dependent objects were found."}`,
		),
	)
	_, e := netbox_tester.New(t, s.URL).Sites()
	assert.True(t, conflict.Is(e))
}

func TestPermissionTokenLacksIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"detail":"You do not have permission to perform this action."}`,
		),
	)
	_, e := netbox_tester.New(t, s.URL).Sites()
	assert.True(t, unexpected.Is(e))
}

func TestBadGatewayInFrontOfNetboxIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadGateway),
		upstream_server.WithAnswer(
			"<html><body><h1>502 Bad Gateway</h1></body></html>",
		),
	)
	_, e := netbox_tester.New(t, s.URL).Sites()
	assert.True(t, unavailable.Is(e))
}
