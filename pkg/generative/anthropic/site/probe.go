package site

import "github.com/funtimecoding/soil/pkg/console"

func (s *Site) Probe() {
	n := s.protocol.Select("div[role='meter']", 0)

	if n == nil {
		console.Line("no meter found")

		return
	}

	console.Format("aria-valuenow: %s\n\n", n.AttributeValue("aria-valuenow"))
	console.Line("--- great-grandparent ---")
	console.Line(s.protocol.Outer("div:has(> div > div > div[role='meter'])"))
}
