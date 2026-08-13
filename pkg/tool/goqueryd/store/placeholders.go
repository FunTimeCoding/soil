package store

import "strings"

func placeholders(count int) string {
	return strings.Join(strings.Split(strings.Repeat("?", count), ""), ", ")
}
