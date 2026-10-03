package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
	"slices"
)

func (s *Site) OpenChat(name string) {
	var names []string
	errors.PanicOnError(
		s.session.Evaluate(
			`Array.from(document.querySelectorAll('a[href^="/c/"] span[dir="auto"]')).map(span => span.textContent.trim())`,
			&names,
		),
	)

	if !slices.Contains(names, name) {
		log.Panicf("chat not found: %s", name)
	}

	s.session.ClickSearch(
		fmt.Sprintf(
			`//a[contains(@href, "/c/")]//span[@dir="auto" and text()="%s"]`,
			name,
		),
	)
}
