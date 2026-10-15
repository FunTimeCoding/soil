package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/option"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
)

func target(o *option.Build) string {
	if o.Module != "" {
		return fmt.Sprintf(
			"%s/%s/%s",
			o.Module,
			systemConstant.CommandPath,
			o.Name,
		)
	}

	return Package(o.MainPath)
}
