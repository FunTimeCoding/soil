package gobundle

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/macos"
	"github.com/funtimecoding/soil/pkg/tool/gobundle/option"
)

func Run(o *option.Bundle) {
	fmt.Printf("Name: %s\n", o.Name)
	fmt.Printf("Path: %s\n", o.Path)
	fmt.Printf("Executable: %s\n", o.Executable)
	fmt.Printf("Icon: %s\n", o.Icon)
	fmt.Printf("Vendor: %s\n", o.Vendor)
	fmt.Printf("Version: %s\n", o.BundleVersion)
	macos.CreateBundle(
		o.Name,
		o.Path,
		o.Executable,
		o.Icon,
		o.Vendor,
		o.BundleVersion,
	)
}
