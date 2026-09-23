package reader

import "github.com/funtimecoding/soil/pkg/console"

func (s *Reader) Probe() {
	n := s.Protocol.Select("div[role='meter']", 0)

	if n == nil {
		console.Line("no meter found")

		return
	}

	console.Format("aria-valuenow: %s\n\n", n.AttributeValue("aria-valuenow"))
	console.Line("--- great-grandparent ---")
	console.Line(s.Protocol.Outer("div:has(> div > div > div[role='meter'])"))
}
