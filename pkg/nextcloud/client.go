package nextcloud

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/basic"
	"github.com/funtimecoding/soil/pkg/web/author"
)

type Client struct {
	basic  *basic.Client
	author *author.Client
}
