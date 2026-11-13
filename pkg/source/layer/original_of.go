package layer

func (l *Layer) originalOf(path string) []byte {
	for _, c := range l.stack() {
		if original, held := c.original[path]; held {
			return original
		}
	}

	return nil
}
