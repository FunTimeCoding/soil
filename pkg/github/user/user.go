package user

import "github.com/google/go-github/v89/github"

type User struct {
	Name string
	Raw  *github.User
}
