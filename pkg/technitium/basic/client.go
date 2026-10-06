package basic

import "github.com/funtimecoding/soil/pkg/web/requester"

type Client struct {
	requester *requester.Requester
	base      string
}
