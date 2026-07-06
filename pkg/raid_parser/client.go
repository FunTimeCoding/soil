package raid_parser

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
