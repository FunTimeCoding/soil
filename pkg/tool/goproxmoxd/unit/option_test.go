package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"testing"
)

func TestTheOptionCarriesBothFactories(t *testing.T) {
	o := option.New()
	assert.NotNil(t, o.Hypervisor)
	assert.NotNil(t, o.Shell)
}

func TestAHypervisorWithoutCredentialsIsRefused(t *testing.T) {
	alfa := instance.Instance{Name: "alfa", Host: "alfa.example"}
	_, e := option.BuildHypervisor(&alfa)
	assert.True(t, validation.Is(e))
	bravo := instance.Instance{Name: "bravo", Token: "bravo", Secret: "secret"}
	_, f := option.BuildHypervisor(&bravo)
	assert.True(t, validation.Is(f))
}

func TestAHypervisorWithATokenIsBuilt(t *testing.T) {
	alfa := instance.Instance{
		Name:   "alfa",
		Host:   "alfa.example",
		Token:  "alfa",
		Secret: "secret",
	}
	c, e := option.BuildHypervisor(&alfa)
	assert.FatalOnError(t, e)
	defer c.Close()
	assert.NotNil(t, c)
}

func TestAShellWithoutCredentialsIsNotConfigured(t *testing.T) {
	alfa := instance.Instance{Name: "alfa", Host: "alfa.example"}
	_, e := option.BuildShell(&alfa)
	assert.True(t, not_configured.Is(e))
}
