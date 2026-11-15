package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server"
	"testing"
)

func TestStorageStatusFilled(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/nodes/alfa/status", "200 OK", `{"data":{}}`)
	s.AnswerMessage(
		"/nodes/alfa/storage/bravo/status",
		"200 OK",
		`{"data":{"type":"dir","content":"iso,snippets","enabled":1}}`,
	)
	c := s.Client()
	n, e := c.Node("alfa")
	assert.FatalOnError(t, e)
	result, e := c.Storage(n, "bravo")
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", result.Node)
	assert.String(t, "bravo", result.Name)
	assert.String(t, "dir", result.Type)
	assert.String(t, "iso,snippets", result.Content)
}

func TestMissingStorageNotFound(t *testing.T) {
	s := proxmox_server.New(t)
	s.AnswerMessage("/nodes/alfa/status", "200 OK", `{"data":{}}`)
	s.AnswerMessage(
		"/nodes/alfa/storage/bravo/status",
		"500 storage 'bravo' does not exist",
		"",
	)
	c := s.Client()
	n, e := c.Node("alfa")
	assert.FatalOnError(t, e)
	_, e = c.Storage(n, "bravo")
	assert.True(t, not_found.Is(e))
}
