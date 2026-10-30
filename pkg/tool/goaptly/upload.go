package goaptly

import (
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goaptly/constant"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

func upload(
	c *aptly.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   "upload [file]",
		Short: "Upload a deb package and add it to a repository",
		Args:  cobra.ExactArgs(1),
		Run: func(
			o *cobra.Command,
			arguments []string,
		) {
			repository, e := o.Flags().GetString(constant.RepositoryFlag)
			cobra.CheckErr(e)
			force, f := o.Flags().GetBool(constant.ForceFlag)
			cobra.CheckErr(f)
			name := filepath.Base(arguments[0])
			directory := strings.TrimSuffix(name, filepath.Ext(name))
			t.EmitDone(c.Upload(directory, arguments[0]))
			t.EmitDone(c.AddToRepository(repository, directory, force))
		},
	}
	result.Flags().String(
		constant.RepositoryFlag,
		constant.DefaultRepository,
		"Repository to add the package to",
	)
	result.Flags().Bool(
		constant.ForceFlag,
		false,
		"Replace a package of the same name and version",
	)

	return result
}
