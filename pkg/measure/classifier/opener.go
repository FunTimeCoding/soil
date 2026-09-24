package classifier

import (
	"github.com/funtimecoding/soil/pkg/measure/language"
	"strings"
)

func (c *Classifier) opener(rest string) *language.Block {
	for _, b := range c.language.BlockComments {
		if strings.HasPrefix(rest, b.Open) {
			return b
		}
	}

	return nil
}
