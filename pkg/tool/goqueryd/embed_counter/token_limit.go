package embed_counter

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (c *Counter) TokenLimit() int {
	return constant.EmbedWindow - c.specials
}
