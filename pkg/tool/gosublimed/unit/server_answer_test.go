package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"testing"
)

func TestAnUnknownViewIsAnswered404(t *testing.T) {
	r, e := newServer(t, "alfa").GetView(
		context.Background(),
		server.GetViewRequestObject{Id: 99},
	)
	assert.FatalOnError(t, e)
	_, okay := r.(server.GetView404JSONResponse)
	assert.True(t, okay)
}

func TestAnEditWithoutOldStringIsAnswered400(t *testing.T) {
	r, e := newServer(t, "alfa").EditView(
		context.Background(),
		server.EditViewRequestObject{
			Id:   1,
			Body: &server.EditViewJSONRequestBody{NewString: "bravo"},
		},
	)
	assert.FatalOnError(t, e)
	_, okay := r.(server.EditView400JSONResponse)
	assert.True(t, okay)
}

func TestAnAmbiguousEditIsAnswered409(t *testing.T) {
	r, e := newServer(t, "alfa alfa").EditView(
		context.Background(),
		server.EditViewRequestObject{
			Id: 1,
			Body: &server.EditViewJSONRequestBody{
				OldString: "alfa",
				NewString: "bravo",
			},
		},
	)
	assert.FatalOnError(t, e)
	_, okay := r.(server.EditView409JSONResponse)
	assert.True(t, okay)
}
