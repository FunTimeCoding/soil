package basic

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

type Client struct {
	requester *requester.Requester
	old       *requester.Requester
	root      *locator.Locator
	base      *locator.Locator
}
