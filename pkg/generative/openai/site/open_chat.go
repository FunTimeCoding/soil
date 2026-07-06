package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"log"
	"slices"
)

func (s *Site) OpenChat(name string) {
	var names []string
	s.protocol.Evaluate(
		`Array.from(document.querySelectorAll('a[href^="/c/"] span[dir="auto"]')).map(span => span.textContent.trim())`,
		&names,
	)

	if !slices.Contains(names, name) {
		log.Panicf("chat not found: %s", name)
	}

	if false {
		fmt.Printf("Name: %s\n", join.Comma(names))
	}

	s.protocol.ClickSearch(
		fmt.Sprintf(
			`//a[contains(@href, "/c/")]//span[@dir="auto" and text()="%s"]`,
			name,
		),
	)
}
