package gobundle

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/macos"
	"github.com/funtimecoding/soil/pkg/tool/gobundle/option"
)

func Run(o *option.Bundle) {
	console.Format("Name: %s\n", o.Name)
	console.Format("Path: %s\n", o.Path)
	console.Format("Executable: %s\n", o.Executable)
	console.Format("Icon: %s\n", o.Icon)
	console.Format("Vendor: %s\n", o.Vendor)
	console.Format("Version: %s\n", o.BundleVersion)
	macos.CreateBundle(
		o.Name,
		o.Path,
		o.Executable,
		o.Icon,
		o.Vendor,
		o.BundleVersion,
	)
}
