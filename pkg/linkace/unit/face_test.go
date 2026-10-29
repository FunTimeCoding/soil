package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linkace"
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/linkace/mock_client"
	"testing"
)

func TestTheLinkAceClientAndItsMockShareTheFace(t *testing.T) {
	c := linkace.New("linkace.example", "alfa-token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}
