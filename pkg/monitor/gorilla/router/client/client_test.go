package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, net.IP{}, nil))
}
