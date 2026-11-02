package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"testing"
)

func TestTheClientAndItsMockKeepTheSSHFace(t *testing.T) {
	c := ssh.NewWithPassword("alfa", "alfa.example", "secret", false)
	m := mock_client.New()
	clients := []face.SSHClient{c, m}
	m.Close()
	assert.Count(t, 2, clients)
	assert.Integer(t, 1, m.Closed)
}
