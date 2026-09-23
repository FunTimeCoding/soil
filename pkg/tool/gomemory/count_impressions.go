package gomemory

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func countImpressions(r *client.ProfileResponse) int {
	if r.Impressions == nil {
		return 0
	}

	return len(*r.Impressions)
}
