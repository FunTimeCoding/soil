package container

import "github.com/google/go-github/v92/github"

type Container struct {
	Name       string
	Repository string
	Raw        *github.Package
}
