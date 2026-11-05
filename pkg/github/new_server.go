package github

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/google/go-github/v93/github"
)

func NewServer(
	token string,
	address string,
) *Client {
	base := join.Empty(address, "/")

	return build(token, github.WithURLs(&base, &base))
}
