package code

import "github.com/google/go-github/v90/github"

func New(v *github.CodeResult) *Code {
	return &Code{
		Hash: v.GetSHA(),
		Name: v.GetName(),
		Path: v.GetPath(),
		Raw:  v,
	}
}
