package site

import "fmt"

func (s *Site) Probe() {
	n := s.protocol.Select("div[role='meter']", 0)

	if n == nil {
		fmt.Println("no meter found")

		return
	}

	fmt.Printf(
		"aria-valuenow: %s\n\n",
		n.AttributeValue("aria-valuenow"),
	)
	fmt.Println("--- great-grandparent ---")
	fmt.Println(
		s.protocol.Outer(
			"div:has(> div > div > div[role='meter'])",
		),
	)
}
