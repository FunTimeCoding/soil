package validate

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"log"
	"slices"
)

func Contains(
	valid []string,
	s string,
) {
	if !slices.Contains(valid, s) {
		log.Panicf("%s not in %s", s, join.Comma(valid))
	}
}
