package build

import (
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func Architectures(o *option.Build) {
	if o.LinuxAMD64 {
		o.OperatingSystem = constant.Linux
		o.Architecture = constant.AMD64
		Go(o)
	}

	if o.DarwinARM64 {
		o.OperatingSystem = constant.Darwin
		o.Architecture = constant.ARM64
		Go(o)
	}

	if o.DarwinAMD64 {
		o.OperatingSystem = constant.Darwin
		o.Architecture = constant.AMD64
		Go(o)
	}
}
