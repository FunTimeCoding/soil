package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) TriggerServed() {
	s.t.Helper()
	v := s.Served()
	c, e := client.NewClientWithResponses(
		v.Address(webConstant.RootPath),
		client.WithHTTPClient(v.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(s.t, e)
	r, f := c.PostTriggerWithResponse(s.t.Context(), client.TriggerRequest{})
	assert.FatalOnError(s.t, f)
	assert.Integer(s.t, http.StatusAccepted, r.StatusCode())
}
