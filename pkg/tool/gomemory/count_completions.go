package gomemory

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func countCompletions(r *client.ProfileResponse) int {
	if r.Completions == nil {
		return 0
	}

	return len(*r.Completions)
}
