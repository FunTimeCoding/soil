package image

import "github.com/google/go-github/v89/github"

func New(v *github.PackageVersion) *Image {
	var tags []string
	meta, okay := v.GetMetadata()

	if okay && meta.Container != nil {
		tags = meta.Container.Tags
	}

	return &Image{
		Identifier: v.GetID(),
		Digest:     v.GetName(),
		Tags:       tags,
		Create:     v.GetCreatedAt().Time,
		Raw:        v,
	}
}
