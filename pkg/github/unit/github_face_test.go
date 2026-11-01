package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/github/mock_client"
	"testing"
)

func TestTheClientAndItsMockKeepTheFace(t *testing.T) {
	m := mock_client.New()
	clients := []face.Client{github.New("alfa-token"), m}

	for _, c := range clients {
		c.Close()
	}

	assert.Integer(t, 1, m.Closed)
}
