package tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"testing"
)

func (o *Tester) Discovery(t *testing.T) *service.DiscoveryDocument {
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
	var d service.DiscoveryDocument
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&d))

	return &d
}
