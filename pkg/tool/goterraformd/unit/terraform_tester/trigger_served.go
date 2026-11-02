package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (o *Tester) TriggerServed() {
	o.t.Helper()
	c, e := client.NewClientWithResponses(
		o.served.Address(webConstant.RootPath),
		client.WithHTTPClient(o.served.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(o.t, e)
	r, f := c.PostTriggerWithResponse(o.t.Context(), client.TriggerRequest{})
	assert.FatalOnError(o.t, f)
	assert.Integer(o.t, http.StatusAccepted, r.StatusCode())
}
