package tag

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.Tag) *Tag {
	result := &Tag{
		Name:      v.Name,
		Message:   v.Message,
		Target:    v.Target,
		Protected: v.Protected,
		Create:    v.CreatedAt,
		Raw:       v,
	}

	if v.Commit != nil {
		result.Hash = v.Commit.ID
	}

	return result
}
