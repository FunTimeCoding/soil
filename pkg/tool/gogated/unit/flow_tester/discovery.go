package flow_tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/response"
)

func (o *Tester) Discovery() *response.Discovery {
	o.t.Helper()
	r, e := o.client.Get(
		fmt.Sprintf(
			"%s/.well-known/openid-configuration",
			o.server.BaseLocator(),
		),
	)
	assert.FatalOnError(o.t, e)

	defer errors.PanicClose(r.Body)
	assert.Integer(o.t, 200, r.StatusCode)
	var d response.Discovery
	assert.FatalOnError(o.t, json.NewDecoder(r.Body).Decode(&d))

	return &d
}
