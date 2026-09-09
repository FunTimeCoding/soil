package registry_repository

import "gitlab.com/gitlab-org/api/client-go/v2"

func NewSlice(v []*gitlab.RegistryRepository) []*Repository {
	result := make([]*Repository, 0, len(v))

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
