package face

import (
	"github.com/funtimecoding/soil/pkg/github/code"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/google/go-github/v93/github"
)

type RepositoryClient interface {
	MustRepositories(owner string) []*github.Repository
	MustSearchRepository(
		query string,
		a ...any,
	) []*repository.Repository
	Repositories(owner string) ([]*github.Repository, error)
	SearchCode(
		query string,
		a ...any,
	) ([]*code.Code, error)
	SearchRepository(
		query string,
		a ...any,
	) ([]*repository.Repository, error)
}
