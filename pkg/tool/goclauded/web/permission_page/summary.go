package permission_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
)

func Summary(r *permission.Report) []string {
	return []string{
		fmt.Sprintf("last %d days", r.Days),
		fmt.Sprintf("%d sessions", r.Sessions),
		fmt.Sprintf("%d prompts", r.Prompts),
		fmt.Sprintf("%d declined", r.Declined),
		fmt.Sprintf("%d the policy would clear", r.Cleared),
	}
}
