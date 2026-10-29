package gohabitica

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/face"
	"github.com/spf13/cobra"
)

func equip(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "equip <key>",
		Short: "Equip a gear item by key",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(c.EquipGear(context.Background(), arguments[0]))
		},
	}
}
