package classifier

import "strings"

func (c *Classifier) quoteOpener(rest string) (string, bool) {
	for _, q := range c.language.RawQuotes {
		if strings.HasPrefix(rest, q) {
			return q, true
		}
	}

	for _, q := range c.language.Quotes {
		if strings.HasPrefix(rest, q) {
			return q, false
		}
	}

	return "", false
}
