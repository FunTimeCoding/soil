package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func handleDialog(w http.ResponseWriter, e *http.Request) {
	c := request_context.New(w, e)
	c.SetLastLocation()
	l := multi_line.New()
	request_context.AppendDetail(l, c)
	r := &model.PostActionIntegrationRequest{}
	notation.MustDecode(c.Body(), r, false)
	l.Format("Decoded request: %+v", r)
	response := &model.PostActionIntegrationResponse{
		EphemeralText: "You updated the post!",
	}

	if false {
		response.Update = &model.Post{Message: "Updated message text!"}
	}

	c.WriteOkay(notation.Encode(response, false))
	console.Line("Request receive:")
	console.Line(l.Render())
}
