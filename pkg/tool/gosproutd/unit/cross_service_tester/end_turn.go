package cross_service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (o *Tester) EndTurn(identifier string) {
	o.t.Helper()
	response, e := o.Goclauded.RESTClient(o.t).PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: identifier},
	)
	assert.FatalOnError(o.t, e)
	assert.Integer(o.t, 200, response.StatusCode())
}
