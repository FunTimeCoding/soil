package channel

import "github.com/funtimecoding/soil/pkg/notation"

func (c *Channel) Encode() []byte {
	return notation.Marshal(c)
}
