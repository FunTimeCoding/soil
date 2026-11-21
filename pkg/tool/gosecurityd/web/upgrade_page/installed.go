package upgrade_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
)

func installed(g *upgrade_group.Group) string {
	if g.InstalledLowest == g.InstalledHighest {
		return g.InstalledLowest
	}

	return fmt.Sprintf("%s – %s", g.InstalledLowest, g.InstalledHighest)
}
