package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	generated "github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/server"
	"testing"
)

func TestViewLifecycle(t *testing.T) {
	s := server.New(mock_client.New())
	created, e := s.CreateView(
		context.Background(),
		generated.CreateViewRequestObject{
			Body: &generated.CreateViewJSONRequestBody{
				Title:   "notes",
				Content: "alfa",
			},
		},
	)
	assert.Nil(t, e)
	v := generated.View(created.(generated.CreateView200JSONResponse))
	assert.Integer(t, 1, v.ViewId)
	assert.String(t, "notes", v.Title)
	list, e := s.GetViews(
		context.Background(),
		generated.GetViewsRequestObject{},
	)
	assert.Nil(t, e)
	assert.Integer(t, 1, len(list.(generated.GetViews200JSONResponse)))
	edited, e := s.EditView(
		context.Background(),
		generated.EditViewRequestObject{
			Id: 1,
			Body: &generated.EditViewJSONRequestBody{
				OldString: "alfa",
				NewString: "bravo",
			},
		},
	)
	assert.Nil(t, e)
	text := generated.View(edited.(generated.EditView200JSONResponse)).Text
	assert.NotNil(t, text)
	assert.String(t, "bravo", *text)
	read, e := s.GetView(
		context.Background(),
		generated.GetViewRequestObject{Id: 1},
	)
	assert.Nil(t, e)
	assert.Integer(
		t,
		1,
		generated.View(read.(generated.GetView200JSONResponse)).ViewId,
	)
	_, e = s.CloseView(
		context.Background(),
		generated.CloseViewRequestObject{Id: 1},
	)
	assert.Nil(t, e)
	list, e = s.GetViews(
		context.Background(),
		generated.GetViewsRequestObject{},
	)
	assert.Nil(t, e)
	assert.Integer(t, 0, len(list.(generated.GetViews200JSONResponse)))
}
