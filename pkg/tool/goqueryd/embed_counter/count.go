package embed_counter

import (
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/tokenizer"
)

func (c *Counter) Count(text string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result, e := c.tokenizer.Encode(
		tokenizer.Valid(text),
		tokenizers.WithAddSpecialTokens(),
	)
	errors.PanicOnError(e)

	return len(result.IDs) - c.specials
}
