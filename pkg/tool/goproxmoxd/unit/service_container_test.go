package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestStartContainerMissing(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	_, e := service_tester.ContainerService(t, c).StartContainer(
		"test",
		999,
		"",
	)
	assert.True(t, not_found.Is(e))
}

func TestStopContainerMissing(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	_, e := service_tester.ContainerService(t, c).StopContainer("test", 999, "")
	assert.True(t, not_found.Is(e))
}

func TestShutdownContainerMissing(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	_, e := service_tester.ContainerService(t, c).ShutdownContainer(
		"test",
		999,
		"",
	)
	assert.True(t, not_found.Is(e))
}

func TestStartContainerFound(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	c.AddContainer("test", 101, "target-container")
	taskIdentifier, e := service_tester.ContainerService(t, c).StartContainer(
		"test",
		101,
		"test",
	)
	assert.Nil(t, e)
	assert.StringContains(t, "ct-start", taskIdentifier)
}
