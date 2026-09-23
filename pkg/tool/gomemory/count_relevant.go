package gomemory

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func countRelevant(r *client.ProfileResponse) int {
	if r.Relevant == nil {
		return 0
	}

	return len(*r.Relevant)
}
