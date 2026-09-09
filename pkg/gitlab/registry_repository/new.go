package registry_repository

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.RegistryRepository) *Repository {
	return &Repository{
		Identifier:        v.ID,
		ProjectIdentifier: v.ProjectID,
		Name:              v.Name,
		Path:              v.Path,
		Location:          v.Location,
		Create:            v.CreatedAt,
		TagCount:          v.TagsCount,
		Raw:               v,
	}
}
