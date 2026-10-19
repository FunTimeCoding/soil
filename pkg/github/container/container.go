package container

import "github.com/google/go-github/v93/github"

type Container struct {
	Name       string
	Repository string
	Raw        *github.Package
}
