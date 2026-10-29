package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	c := client.New(locator.New("goatlasd.test"), "token")
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}
