package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestTheDirectoryClientAndItsMockShareTheFace(t *testing.T) {
	faces := []face.Client{
		directory.New("directory.test", "dc=test"),
		mock_client.New("dc=test"),
	}
	assert.Count(t, 2, faces)
}

func TestARefusedDialIsAnOutage(t *testing.T) {
	c := directory.New("127.0.0.1", "dc=test").
		WithInsecure().
		WithPort(system.FindUnusedPort(19500)).
		WithCredential("cn=admin,dc=test", "secret")
	e := c.Reachable(t.Context())
	assert.NotNil(t, connection.Classify(e))
	assert.True(t, unreachable.Is(e))
	assert.Error(t, c.Add("uid=alfa,dc=test", map[string][]string{}))
}
