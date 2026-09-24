package gohook

import (
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/spf13/cobra"
)

func uninstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constant.Uninstall,
		Short: "Remove hook stubs gohook wrote",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			Uninstall(configuration.Load(root()))
		},
	}
}
