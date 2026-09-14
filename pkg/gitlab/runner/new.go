package runner

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.Runner) *Runner {
	return &Runner{
		Identifier:  v.ID,
		Name:        v.Name,
		Description: v.Description,
		Status:      v.Status,
		Type:        v.RunnerType,
		Online:      v.Online,
		Paused:      v.Paused,
		Shared:      v.IsShared,
		RawList:     v,
	}
}
