package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestClientAndItsMockKeepFace(t *testing.T) {
	r := client.New(locator.New("localhost").Insecure(), "india")
	m := mock_client.New()
	clients := []face.Client{r, m}

	for _, c := range clients {
		c.Close()
	}

	assert.Integer(t, 1, m.Closed)
}
