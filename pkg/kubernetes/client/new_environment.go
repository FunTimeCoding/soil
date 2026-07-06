package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	var contexts []string

	if s := environment.Optional(constant.ContextEnvironment); s != "" {
		contexts = split.Comma(s)
	}

	return New(contexts)
}
