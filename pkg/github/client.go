package github

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/google/go-github/v92/github"
)

type Client struct {
	client  *github.Client
	context context.Context
	pool    *pool.Pool
}
