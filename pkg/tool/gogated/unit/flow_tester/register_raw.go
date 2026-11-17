package flow_tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (o *Tester) RegisterRaw(body map[string]any) int {
	o.t.Helper()
	encoded, e := json.Marshal(body)
	assert.FatalOnError(o.t, e)
	r, e := o.client.Post(
		fmt.Sprintf("%s/register", o.server.BaseLocator()),
		"application/json",
		bytes.NewReader(encoded),
	)
	assert.FatalOnError(o.t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
