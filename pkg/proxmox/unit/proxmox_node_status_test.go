package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server"
	"testing"
)

func TestNodeStatusReadsFractionalWait(t *testing.T) {
	r := proxmox_server.New(t)
	r.AnswerMessage(
		"/nodes/alfa/status",
		"200 OK",
		`{"data":{"cpu":0.05,"wait":0.000123456789012345,"idle":0}}`,
	)
	s, e := r.Client().NodeStatus("alfa")
	assert.FatalOnError(t, e)
	assert.Float(t, 0.000123456789012345, s.Wait)
}
