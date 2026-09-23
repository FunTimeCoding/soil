package site

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/reader"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func Attach(c *chromium.Client) *reader.Reader {
	return reader.New(protocol.New(c, constant.AnthropicUsageFragment))
}
