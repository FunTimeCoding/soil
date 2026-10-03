package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	stampConstant "github.com/funtimecoding/soil/pkg/stamp/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func LinkerFlags(
	version string,
	hash string,
	date string,
	module string,
	dirty bool,
) string {
	d := stringsConstant.BooleanFalse

	if dirty {
		d = stringsConstant.BooleanTrue
	}

	return join.Space(
		constant.LinkerSetVariable,
		fmt.Sprintf("main.Version=%s", version),
		constant.LinkerSetVariable,
		fmt.Sprintf("main.GitHash=%s", hash),
		constant.LinkerSetVariable,
		fmt.Sprintf("main.BuildDate=%s", date),
		constant.LinkerSetVariable,
		fmt.Sprintf("%s=%s", stampConstant.ModuleVariable, module),
		constant.LinkerSetVariable,
		fmt.Sprintf("%s=%s", stampConstant.DirtyVariable, d),
	)
}
