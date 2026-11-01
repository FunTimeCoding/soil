package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestCloseClosesBothPeerClientsOnce(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Close()
	s.Service.Close()
	assert.Integer(t, 1, s.Memory.Closed)
	assert.Integer(t, 1, s.Query.Closed)
}
