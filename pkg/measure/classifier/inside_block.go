package classifier

import "strings"

func (c *Classifier) insideBlock(rest string) int {
	last := c.open[len(c.open)-1]

	if strings.HasPrefix(rest, last.Close) {
		c.open = c.open[:len(c.open)-1]

		return len(last.Close)
	}

	if c.language.Nested && strings.HasPrefix(rest, last.Open) {
		c.open = append(c.open, last)

		return len(last.Open)
	}

	return 1
}
