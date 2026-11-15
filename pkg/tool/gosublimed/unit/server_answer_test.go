package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/client_tester"
	"net/http"
	"testing"
)

func TestUnknownViewIsAnswered404(t *testing.T) {
	r, e := client_tester.ServedView(t, "alfa").GetViewWithResponse(
		context.Background(),
		99,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.NotNil(t, r.JSON404)
}

func TestEditWithoutOldStringIsAnswered400(t *testing.T) {
	r, e := client_tester.ServedView(t, "alfa").EditViewWithResponse(
		context.Background(),
		1,
		client.EditViewJSONRequestBody{NewString: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
}

func TestAmbiguousEditIsAnswered409(t *testing.T) {
	r, e := client_tester.ServedView(t, "alfa alfa").EditViewWithResponse(
		context.Background(),
		1,
		client.EditViewJSONRequestBody{OldString: "alfa", NewString: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusConflict, r.StatusCode())
}
