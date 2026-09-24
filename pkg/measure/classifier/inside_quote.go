package classifier

import "strings"

func (c *Classifier) insideQuote(rest string) int {
	if !c.raw && rest[0] == '\\' && len(rest) > 1 {
		return 2
	}

	if strings.HasPrefix(rest, c.quote) {
		width := len(c.quote)
		c.quote = ""
		c.raw = false

		return width
	}

	return 1
}
