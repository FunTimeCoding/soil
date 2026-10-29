package face

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/google/go-github/v93/github"
)

type RunSource interface {
	Repositories(owner string) ([]*github.Repository, error)
	LatestRuns(
		owner string,
		repo string,
	) ([]*run.Run, error)
}
