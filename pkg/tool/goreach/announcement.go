package goreach

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/reach"
)

func Announcement(
	r *reach.Reach,
	edits *reach.Edits,
	files []string,
) string {
	return fmt.Sprintf(
		"goreach: %d changed files, %d package directories reached; Bash writes %d resolved, %d globbed, %d unresolved, %d fixer runs, %d scripts",
		len(files),
		len(r.Packages),
		edits.Resolved,
		edits.Globbed,
		edits.Unresolved,
		edits.Fixers,
		edits.Scripts,
	)
}
