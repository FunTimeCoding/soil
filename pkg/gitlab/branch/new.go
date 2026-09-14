package branch

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.Branch) *Branch {
	return &Branch{
		Name:       v.Name,
		Merged:     v.Merged,
		CommitDate: v.Commit.CreatedAt,
		Raw:        v,
	}
}
