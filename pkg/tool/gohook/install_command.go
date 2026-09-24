package gohook

import (
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/spf13/cobra"
)

func installCommand() *cobra.Command {
	return &cobra.Command{
		Use:   constant.Install,
		Short: "Write hook stubs for every hook the configuration names",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			Install(configuration.Load(root()))
		},
	}
}
