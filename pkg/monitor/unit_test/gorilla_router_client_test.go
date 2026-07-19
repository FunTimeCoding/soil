package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, client.New(upper.Alfa, upper.Bravo, net.IP{}, nil))
}
