package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

type Client struct {
	base  *locator.Locator
	token string
}
