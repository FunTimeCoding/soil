package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"testing"
)

func TestTheSaltClientAndItsMockShareTheFace(t *testing.T) {
	c := salt.New("salt.test", 1, "alfa", "bravo")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.SaltClient{c, m}
	assert.Count(t, 2, faces)
}
