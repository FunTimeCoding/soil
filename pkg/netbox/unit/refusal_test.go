package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestARejectedWriteIsValidation(t *testing.T) {
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

func TestAMissingObjectIsNotFound(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusNotFound,
		`{"detail":"No Site matches the given query."}`,
	).Sites()
	assert.True(t, not_found.Is(e))
	assert.String(t, "No Site matches the given query.", e.Error())
}

func TestAConflictIsConflict(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusConflict,
		`{"detail":"Unable to delete: dependent objects were found."}`,
	).Sites()
	assert.True(t, conflict.Is(e))
}

func TestARefusedTokenIsForbidden(t *testing.T) {
	_, e := refusingNetbox(
		t,
		http.StatusForbidden,
		`{"detail":"You do not have permission to perform this action."}`,
	).Sites()
	assert.True(t, forbidden.Is(e))
}
