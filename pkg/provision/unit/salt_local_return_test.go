package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"testing"
)

func TestLocalReturnSilentMinionDecodesAsBool(t *testing.T) {
	var r response.Local
	assert.FatalOnError(
		t,
		json.Unmarshal(
			[]byte(
				`{"return":[{"quiet":false,"loud":{"jid":"1","ret":true,"retcode":0}}]}`,
			),
			&r,
		),
	)
	assert.Count(t, 1, r.Return)
	quiet := r.Return[0]["quiet"]
	loud := r.Return[0]["loud"]
	assert.False(t, quiet.Responded)
	assert.True(t, loud.Responded)
	assert.String(t, "1", loud.Jid)
	assert.Integer(t, 0, loud.Retcode)
}
