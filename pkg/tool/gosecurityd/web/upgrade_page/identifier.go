package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
)

func identifier(g *upgrade_group.Group) string {
	return constant.RowUnsafePattern.ReplaceAllString(
		join.Empty(constant.UpgradeRowPrefix, g.Release, "-", g.Source),
		"-",
	)
}
