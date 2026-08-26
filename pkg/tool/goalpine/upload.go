package goalpine

import (
	"fmt"
	alpine "github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpine/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/client"
	"github.com/spf13/cobra"
)

func upload(c *client.Client) *cobra.Command {
	result := &cobra.Command{
		Use:   "upload [file]",
		Short: "Upload an apk package and rebuild the index",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			arguments []string,
		) {
			version, e := o.Flags().GetString(constant.VersionFlag)
			cobra.CheckErr(e)
			repository, f := o.Flags().GetString(constant.RepositoryFlag)
			cobra.CheckErr(f)
			architecture, g := o.Flags().GetString(constant.ArchitectureFlag)
			cobra.CheckErr(g)
			fmt.Print(c.Upload(arguments[0], version, repository, architecture))
		},
	}
	result.Flags().String(
		constant.VersionFlag,
		constant.DefaultVersion,
		"Repository version",
	)
	result.Flags().String(
		constant.RepositoryFlag,
		constant.DefaultRepository,
		"Repository component",
	)
	result.Flags().String(
		constant.ArchitectureFlag,
		alpine.Architecture,
		"Package architecture",
	)

	return result
}
