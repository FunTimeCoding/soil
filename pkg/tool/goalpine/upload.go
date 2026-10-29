package goalpine

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goalpine/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func upload(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
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
			body, h := os.Open(arguments[0])
			cobra.CheckErr(h)
			defer errors.LogClose(body)
			t.EmitDone(
				c.PostPackage(
					context.Background(),
					version,
					repository,
					architecture,
					filepath.Base(arguments[0]),
					body,
				),
			)
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
		systemConstant.AlpineArchitecture,
		"Package architecture",
	)

	return result
}
