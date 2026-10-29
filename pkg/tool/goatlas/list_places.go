package goatlas

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/spf13/cobra"
)

func listPlaces(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-places",
		Short: "List devices and virtual machines by how much they carry",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.EmitValue(c.ListPlaces(context.Background()))
		},
	}
}
