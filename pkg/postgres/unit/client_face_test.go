package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/postgres"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/postgres/mock_client"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	c, e := postgres.New("postgres://alfa:bravo@127.0.0.1:1/charlie")
	assert.FatalOnError(t, e)
	t.Cleanup(c.Close)
	m := mock_client.New()
	t.Cleanup(m.Close)
	assert.Count(t, 2, []face.Client{c, m})
}
