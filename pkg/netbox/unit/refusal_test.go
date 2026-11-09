package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestRejectedWriteIsValidation(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusBadRequest,
		`{"slug":["site with this slug already exists."],"name":["required"]}`,
	).CreateSite("alfa")
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"name: required; slug: site with this slug already exists.",
		e.Error(),
	)
}

func TestMissingObjectIsNotFound(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusNotFound,
		`{"detail":"No Site matches the given query."}`,
	).Sites()
	assert.True(t, not_found.Is(e))
	assert.String(t, "No Site matches the given query.", e.Error())
}

func TestConflictIsConflict(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusConflict,
		`{"detail":"Unable to delete: dependent objects were found."}`,
	).Sites()
	assert.True(t, conflict.Is(e))
}

func TestPermissionTokenLacksIsUnexpected(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusForbidden,
		`{"detail":"You do not have permission to perform this action."}`,
	).Sites()
	assert.True(t, unexpected.Is(e))
}

func TestBadGatewayInFrontOfNetboxIsUnavailable(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusBadGateway,
		"<html><body><h1>502 Bad Gateway</h1></body></html>",
	).Sites()
	assert.True(t, unavailable.Is(e))
}
