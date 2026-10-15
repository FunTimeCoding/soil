package gobuild

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
	"strings"
)

func Validate(o *option.Build) error {
	for _, n := range o.Names {
		if n == "" || strings.ContainsAny(n, ",/ ") {
			return fmt.Errorf(
				"invalid name %q: one program name per argument",
				n,
			)
		}
	}

	if o.Module != "" {
		if len(o.Names) == 0 {
			return fmt.Errorf("--%s needs program names", constant.ModuleFlag)
		}

		if o.Main != "" || o.Output != "" {
			return fmt.Errorf(
				"--%s and --%s do not apply to --%s",
				constant.MainFlag,
				constant.OutputFlag,
				constant.ModuleFlag,
			)
		}

		path, version, _ := strings.Cut(o.Module, constant.VersionSeparator)

		if path == "" || strings.HasSuffix(o.Module, constant.VersionSeparator) {
			return fmt.Errorf("invalid module %q: path[@version]", o.Module)
		}

		o.Module = path
		o.Version = version
	}

	if o.Main != "" && len(o.Names) != 1 {
		return fmt.Errorf("--%s needs exactly one name", constant.MainFlag)
	}

	if o.Output != "" && (len(o.Names) != 1 || len(o.Targets) != 1) {
		return fmt.Errorf(
			"--%s needs exactly one name and one target",
			constant.OutputFlag,
		)
	}

	return nil
}
