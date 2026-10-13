package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"testing"
)

func TestInventoryRefusesAnUnscopedPublicHost(t *testing.T) {
	e := inventory.New[instance.Instance](
		instance.Instance{Name: "bravo", Host: "https://gitlab.com"},
	).Validate()
	assert.True(t, validation.Is(e))
	assert.String(
		t,
		"instance bravo points at gitlab.com without groups or projects",
		e.Error(),
	)
}

func TestInventoryAllowsAScopedPublicHost(t *testing.T) {
	assert.FatalOnError(
		t,
		inventory.New[instance.Instance](
			instance.Instance{
				Name:     "bravo",
				Host:     "gitlab.com",
				Projects: []int64{42},
			},
		).Validate(),
	)
}

func TestInventoryAllowsAnUnscopedSelfHostedInstance(t *testing.T) {
	assert.FatalOnError(
		t,
		inventory.New[instance.Instance](
			instance.Instance{Name: "alfa", Host: "gitlab.example"},
		).Validate(),
	)
}
