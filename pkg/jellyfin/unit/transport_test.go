package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/jellyfin/transport"
	"testing"
)

func TestTransport(t *testing.T) {
	assert.NotNil(t, transport.New(""))
}
