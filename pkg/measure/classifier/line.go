package classifier

func (c *Classifier) line(trimmed string) bool {
	if len(c.open) == 0 && c.quote == "" && c.isLineComment(trimmed) {
		return true
	}

	code := false
	position := 0

	for position < len(trimmed) {
		rest := trimmed[position:]

		switch {
		case len(c.open) > 0:
			position += c.insideBlock(rest)
		case c.quote != "":
			code = true
			position += c.insideQuote(rest)
		default:
			if b := c.opener(rest); b != nil {
				c.open = append(c.open, b)
				position += len(b.Open)

				continue
			}

			if q, raw := c.quoteOpener(rest); q != "" {
				c.quote = q
				c.raw = raw
				code = true
				position += len(q)

				continue
			}

			if !isSpace(trimmed[position]) {
				code = true
			}

			position++
		}
	}

	return !code
}
