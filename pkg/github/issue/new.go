package issue

import (
	"github.com/google/go-github/v89/github"
	"net/url"
	"strings"
)

func New(v *github.Issue) *Issue {
	var repository string

	if u, e := url.Parse(*v.RepositoryURL); e == nil {
		repository = strings.TrimPrefix(u.Path, "/repos/")
	}

	return &Issue{
		Repository: repository,
		Title:      *v.Title,
		Link:       *v.HTMLURL,
		Raw:        v,
	}
}
