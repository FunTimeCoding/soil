package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/mock_client"
	"reflect"
	"testing"
)

func TestTheChromiumClientAndItsMockShareTheFace(t *testing.T) {
	f := reflect.TypeFor[face.ChromiumClient]()
	assert.True(t, reflect.TypeFor[*chromium.Client]().Implements(f))
	assert.True(t, reflect.TypeFor[*mock_client.Client]().Implements(f))
}

func TestTheAddressJoinsHostAndPort(t *testing.T) {
	c := chromium.New("alfa.example", 9222)
	defer c.Close()
	assert.String(t, "alfa.example:9222", c.Address())
}
