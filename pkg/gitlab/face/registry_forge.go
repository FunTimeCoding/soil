package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type RegistryForge interface {
	RegistryRepositories(
		project int64,
		panicOnForbidden bool,
	) ([]*registry_repository.Repository, error)
	DeleteRegistryRepository(
		project int64,
		repository int64,
	) error
	Images(
		project int64,
		repository int64,
	) ([]*image.Image, error)
	DeleteImage(
		project int64,
		repository int64,
		tag string,
	) error
	DeletePackage(
		project int64,
		repository int64,
	) error
	MustDeleteImage(
		project int64,
		repository int64,
		tag string,
	)
	MustDeletePackage(
		project int64,
		repository int64,
	)
	MustImages(
		project int64,
		repository int64,
	) []*image.Image
	MustPackageFiles(
		project int64,
		packageIdentifier int64,
	) []*gitlab.PackageFile
	MustPackages(
		project int64,
		panicOnForbidden bool,
	) []*gitlab.Package
	MustRegistryRepositories(
		project int64,
		panicOnForbidden bool,
	) []*registry_repository.Repository
	PackageFiles(
		project int64,
		packageIdentifier int64,
	) ([]*gitlab.PackageFile, error)
	Packages(
		project int64,
		panicOnForbidden bool,
	) ([]*gitlab.Package, error)
}
