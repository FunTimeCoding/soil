package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkEntryGuards(
	all []*packages.Package,
	p *packages.Package,
	entries []*relocation.Entry,
	packagePath string,
	targetPackagePath string,
) string {
	if cycle := cyclePath(
		postMoveEdges(all, p, entries, packagePath, targetPackagePath),
		targetPackagePath,
	); cycle != nil {
		return fmt.Sprintf(
			"move would create an import cycle: %s",
			strings.Join(cycle, " → "),
		)
	}

	return ""
}
