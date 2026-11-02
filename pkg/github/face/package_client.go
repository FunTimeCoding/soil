package face

import (
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/github/image"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/google/go-github/v93/github"
)

type PackageClient interface {
	DeletePackageVersion(
		packageName string,
		versionIdentifier int64,
	) error
	DeleteTag(
		owner string,
		repository string,
		name string,
	) error
	MustDeletePackageVersion(
		packageName string,
		versionIdentifier int64,
	)
	MustDeleteTag(
		owner string,
		repository string,
		name string,
	)
	MustPackageVersions(packageName string) []*image.Image
	MustPackages(user string) []*container.Container
	MustReleases(
		namespace string,
		repository string,
	) []*release.Release
	MustTags(
		owner string,
		repository string,
	) []*github.RepositoryTag
	PackageVersions(packageName string) ([]*image.Image, error)
	Packages(user string) ([]*container.Container, error)
	Releases(
		namespace string,
		repository string,
	) ([]*release.Release, error)
	Tags(
		owner string,
		repository string,
	) ([]*github.RepositoryTag, error)
}
