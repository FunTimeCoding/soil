package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/unit/instance_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestInventoryNamesKeepTheFileOrder(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		instance_tester.Inventory("alfa", "bravo").Names(),
	)
}

func TestInventoryFindsAnInstanceByName(t *testing.T) {
	i, okay := instance_tester.Inventory("alfa", "bravo").Find("bravo")
	assert.True(t, okay)
	assert.String(t, "bravo.example", i.Host)
	_, missing := instance_tester.Inventory("alfa").Find("charlie")
	assert.False(t, missing)
}

func TestInventoryResolveInstanceDefaultsToTheOnlyOne(t *testing.T) {
	name, e := instance_tester.Inventory("alfa").ResolveInstance("")
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", name)
}

func TestInventoryResolveInstanceRequiresSelectionWithTwo(t *testing.T) {
	_, e := instance_tester.Inventory("alfa", "bravo").ResolveInstance("")
	assert.True(t, not_selected.Is(e))
}

func TestInventoryResolveInstanceRefusesAnUnknownName(t *testing.T) {
	_, e := instance_tester.Inventory("alfa").ResolveInstance("charlie")
	assert.True(t, not_found.Is(e))
}

func TestInventoryValidateRefusesABrokenInstance(t *testing.T) {
	broken := instance_tester.New("bravo", "bravo.example")
	broken.Broken = true
	e := inventory.New[instance_tester.Instance](
		*instance_tester.New("alfa", "alfa.example"),
		*broken,
	).Validate()
	assert.True(t, validation.Is(e))
	assert.String(t, "instance bravo is broken", e.Error())
}

func TestInventoryLoadReadsTheFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "inventory.yaml")
	assert.FatalOnError(
		t,
		os.WriteFile(
			path,
			[]byte("instances:\n- name: alfa\n  host: alfa.example\n"),
			0o600,
		),
	)
	i := inventory.Load[instance_tester.Instance](path)
	assert.Strings(t, []string{"alfa"}, i.Names())
	assert.String(t, "alfa.example", i.Instances[0].Host)
}
