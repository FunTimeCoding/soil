package confluence

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
)

type Client struct {
	host         string
	labels       []string
	context      context.Context
	basic        *basic.Client
	defaultSpace string
	defaultPage  string
	verbose      bool
}
