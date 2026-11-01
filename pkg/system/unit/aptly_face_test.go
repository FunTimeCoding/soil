package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly/face"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly/mock_client"
	"testing"
)

func TestTheAptlyClientAndItsMockKeepTheFace(t *testing.T) {
	m := mock_client.New()
	clients := []face.Client{
		aptly.New("aptly.example", 8080, true, "alfa", "bravo"),
		m,
	}

	for _, c := range clients {
		c.Close()
	}

	assert.Integer(t, 1, m.Closed)
}
