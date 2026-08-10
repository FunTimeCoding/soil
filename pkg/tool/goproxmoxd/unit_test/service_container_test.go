package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func TestStartContainerMissing(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	s := service.New(inventory.NewSingle("test"))
	_, e := s.StartContainer(c, 999, "")
	assert.True(t, not_found.Is(e))
}

func TestStopContainerMissing(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	s := service.New(inventory.NewSingle("test"))
	_, e := s.StopContainer(c, 999, "")
	assert.True(t, not_found.Is(e))
}

func TestShutdownContainerMissing(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	s := service.New(inventory.NewSingle("test"))
	_, e := s.ShutdownContainer(c, 999, "")
	assert.True(t, not_found.Is(e))
}

func TestStartContainerFound(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddContainer("test", 101, "target-container")
	s := service.New(inventory.NewSingle("test"))
	taskIdentifier, e := s.StartContainer(c, 101, "test")
	assert.True(t, e == nil)
	assert.StringContains(t, "ct-start", taskIdentifier)
}
