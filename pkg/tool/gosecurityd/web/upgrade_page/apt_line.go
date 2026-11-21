package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"slices"
)

func aptLine(binaries []string) string {
	return join.Space(slices.Concat([]string{constant.AptUpgrade}, binaries)...)
}
