package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func placeholders(count int) string {
	return join.CommaSpace(strings.Split(strings.Repeat("?", count), ""))
}
