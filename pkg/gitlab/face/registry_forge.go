package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type RegistryForge interface {
	RegistryRepositories(
		x context.Context,
		project int64,
		panicOnForbidden bool,
	) ([]*registry_repository.Repository, error)
	DeleteRegistryRepository(
		x context.Context,
		project int64,
		repository int64,
	) error
	Images(
		x context.Context,
		project int64,
		repository int64,
	) ([]*image.Image, error)
	DeleteImage(
		x context.Context,
		project int64,
		repository int64,
		tag string,
	) error
	DeletePackage(
		x context.Context,
		project int64,
		repository int64,
	) error
	MustDeleteImage(
		x context.Context,
		project int64,
		repository int64,
		tag string,
	)
	MustDeletePackage(
		x context.Context,
		project int64,
		repository int64,
	)
	MustImages(
		x context.Context,
		project int64,
		repository int64,
	) []*image.Image
	MustPackageFiles(
		x context.Context,
		project int64,
		packageIdentifier int64,
	) []*gitlab.PackageFile
	MustPackages(
		x context.Context,
		project int64,
		panicOnForbidden bool,
	) []*gitlab.Package
	MustRegistryRepositories(
		x context.Context,
		project int64,
		panicOnForbidden bool,
	) []*registry_repository.Repository
	PackageFiles(
		x context.Context,
		project int64,
		packageIdentifier int64,
	) ([]*gitlab.PackageFile, error)
	Packages(
		x context.Context,
		project int64,
		panicOnForbidden bool,
	) ([]*gitlab.Package, error)
}
