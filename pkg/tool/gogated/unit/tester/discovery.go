package tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service/response"
	"testing"
)

func (o *Tester) Discovery(t *testing.T) *response.Discovery {
	t.Helper()
	r, e := o.client.Get(
		fmt.Sprintf(
			"%s/.well-known/openid-configuration",
			o.server.BaseLocator(),
		),
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r.Body)
	assert.Integer(t, 200, r.StatusCode)
	var d response.Discovery
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&d))

	return &d
}
