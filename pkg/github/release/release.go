package release

import (
	"github.com/google/go-github/v89/github"
	"time"
)

type Release struct {
	Name   string
	Create time.Time
	Raw    *github.RepositoryRelease
}
