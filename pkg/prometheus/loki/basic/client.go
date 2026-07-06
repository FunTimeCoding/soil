package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

type Client struct {
	user     string
	password string
	verbose  bool
	base     *locator.Locator
}
