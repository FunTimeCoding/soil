package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func (o *Tester) RegisterRaw(
	t *testing.T,
	body map[string]any,
) int {
	t.Helper()
	encoded, e := json.Marshal(body)
	assert.FatalOnError(t, e)
	r, e := o.client.Post(
		fmt.Sprintf("%s/register", o.server.BaseLocator()),
		"application/json",
		bytes.NewReader(encoded),
	)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
