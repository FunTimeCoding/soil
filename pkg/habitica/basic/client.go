package basic

import "github.com/funtimecoding/soil/pkg/web/requester"

type Client struct {
	requester *requester.Requester
	refusal   func(int, []byte) error
}
