package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"testing"
)

func TestClientAndItsMockKeepBandFace(t *testing.T) {
	c := band.New("http://alfa.example:16992", constant.DefaultUser, "secret")
	t.Cleanup(c.Close)
	m := mock_client.New()
	clients := []face.BandClient{c, m}
	m.Close()
	assert.Count(t, 2, clients)
	assert.Integer(t, 1, m.Closed)
}
