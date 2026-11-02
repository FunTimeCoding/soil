package unit

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
)

func alfaOption() *option.Band {
	result := option.New()
	result.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example"},
	)

	return result
}
