package image

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.RegistryRepositoryTag) *Image {
	return &Image{
		Name:     v.Name,
		Path:     v.Path,
		Location: v.Location,
		Digest:   v.Digest,
		Revision: v.Revision,
		Create:   v.CreatedAt,
		Size:     v.TotalSize,
		Raw:      v,
	}
}
