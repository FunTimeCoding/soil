package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/proxmox/unit/upstream_tester"
	"testing"
)

func TestNodeStatusReadsFractionalWait(t *testing.T) {
	u := upstream_tester.New(t)
	u.AnswerMessage(
		"/nodes/alfa/status",
		"200 OK",
		`{"data":{"cpu":0.05,"wait":0.000123456789012345,"idle":0}}`,
	)
	s, e := u.Client().NodeStatus("alfa")
	assert.FatalOnError(t, e)
	assert.Float(t, 0.000123456789012345, s.Wait)
}
