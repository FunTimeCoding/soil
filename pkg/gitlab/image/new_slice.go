package image

import "gitlab.com/gitlab-org/api/client-go/v3"

func NewSlice(v []*gitlab.RegistryRepositoryTag) []*Image {
	result := make([]*Image, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
