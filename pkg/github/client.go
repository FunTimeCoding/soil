package github

import (
	"context"
	"github.com/google/go-github/v90/github"
)

type Client struct {
	client  *github.Client
	context context.Context
}
