package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"testing"
)

func TestInventoryRefusesUnscopedPublicHost(t *testing.T) {
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

func TestInventoryAllowsScopedPublicHost(t *testing.T) {
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

func TestInventoryAllowsUnscopedSelfHostedInstance(t *testing.T) {
	assert.FatalOnError(
		t,
		inventory.New[instance.Instance](
			instance.Instance{Name: "alfa", Host: "gitlab.example"},
		).Validate(),
	)
}
