package sink

import (
	"bytes"
	"os"
)

func (s *Sink) Stale() []string {
	var result []string

	for _, o := range s.rewrites {
		current, e := os.ReadFile(o.Path)

		if e != nil || !bytes.Equal(current, o.Expected) {
			result = append(result, o.Path)
		}
	}

	return result
}
