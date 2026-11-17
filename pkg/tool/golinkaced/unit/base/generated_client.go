package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (s *Server) GeneratedClient() *client.ClientWithResponses {
	s.t.Helper()
	v := s.Served()
	result, e := client.NewClientWithResponses(
		v.Address(webConstant.RootPath),
		client.WithHTTPClient(v.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(s.t, e)

	return result
}
