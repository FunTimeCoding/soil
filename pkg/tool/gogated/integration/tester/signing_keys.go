package tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"testing"
)

func (o *Tester) SigningKeys(t *testing.T) *service.SigningKeySet {
	t.Helper()
	r, e := o.client.Get(fmt.Sprintf("%s/jwks", o.server.BaseLocator()))
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r.Body)
	assert.Integer(t, 200, r.StatusCode)
	var keys service.SigningKeySet
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&keys))

	return &keys
}
