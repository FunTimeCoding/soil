package classifier

import "strings"

func (c *Classifier) isLineComment(trimmed string) bool {
	if c.opener(trimmed) != nil {
		return false
	}

	for _, prefix := range c.language.LineComments {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}

	return false
}
