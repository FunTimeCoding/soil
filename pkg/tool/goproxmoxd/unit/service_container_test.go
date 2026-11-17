package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestStartContainerMissing(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service().StartContainer("alfa", 999, "")
	assert.True(t, not_found.Is(e))
}

func TestStopContainerMissing(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service().StopContainer("alfa", 999, "")
	assert.True(t, not_found.Is(e))
}

func TestShutdownContainerMissing(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service().ShutdownContainer("alfa", 999, "")
	assert.True(t, not_found.Is(e))
}

func TestStartContainerFound(t *testing.T) {
	o := service_tester.New(t)
	o.Alfa().Hypervisor.AddContainer("alfa", 101, "target-container")
	taskIdentifier, e := o.Service().StartContainer("alfa", 101, "alfa")
	assert.Nil(t, e)
	assert.StringContains(t, "ct-start", taskIdentifier)
}
