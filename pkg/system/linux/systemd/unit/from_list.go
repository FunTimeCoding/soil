package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func FromList(line string) *Unit {
	s := strings.Fields(line)

	return &Unit{
		Name:        s[0],
		Load:        s[1],
		Active:      s[2],
		Sub:         s[3],
		Description: join.Space(s[4:]...),
	}
}
