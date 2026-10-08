package reader

import "github.com/funtimecoding/soil/pkg/console"

func (s *Reader) Probe() {
	n, okay := s.Protocol.MustFindNode("div[role='meter']", 0)

	if !okay {
		console.Line("no meter found")

		return
	}

	console.Format("aria-valuenow: %s\n\n", n.AttributeValue("aria-valuenow"))
	console.Line("--- great-grandparent ---")
	console.Line(
		s.Protocol.MustOuter("div:has(> div > div > div[role='meter'])"),
	)
}
