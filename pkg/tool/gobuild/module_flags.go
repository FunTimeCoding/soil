package gobuild

import (
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
	"github.com/spf13/cobra"
)

func moduleFlags(
	c *cobra.Command,
	o *option.Build,
) {
	f := c.Flags()
	f.StringVar(
		&o.Module,
		constant.ModuleFlag,
		"",
		"Build from a module path[@version] instead of the work directory, latest tag by default",
	)
	f.BoolVar(&o.Native, constant.NativeFlag, false, "Enable CGO")
	f.StringVar(&o.Tags, constant.TagsFlag, "", "Build tags")
}
