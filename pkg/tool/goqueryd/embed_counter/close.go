package embed_counter

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Counter) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.tokenizer != nil {
		errors.PanicOnError(c.tokenizer.Close())
		c.tokenizer = nil
	}

	return nil
}
