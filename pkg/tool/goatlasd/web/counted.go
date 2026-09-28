package web

import "fmt"

func Counted(
	count int,
	singular string,
	plural string,
) string {
	return fmt.Sprintf("%d %s", count, Word(count, singular, plural))
}
