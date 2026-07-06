package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

type Client struct {
	host  string
	user  string
	token string
	base  *locator.Locator
}
