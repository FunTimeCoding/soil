package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func LinkerFlags(
	version string,
	hash string,
	date string,
) string {
	return join.Space(
		constant.LinkerSetVariable,
		fmt.Sprintf("main.Version=%s", version),
		constant.LinkerSetVariable,
		fmt.Sprintf("main.GitHash=%s", hash),
		constant.LinkerSetVariable,
		fmt.Sprintf("main.BuildDate=%s", date),
	)
}
