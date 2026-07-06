package channel

import "github.com/funtimecoding/soil/pkg/notation"

func (c *Channel) Decode(b []byte) {
	notation.MustDecodeBytes(b, c, false)
}
