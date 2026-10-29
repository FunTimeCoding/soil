package openai

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/openai/openai-go/v3"
)

type Client struct {
	client openai.Client
	model  string
	pool   *pool.Pool
}
