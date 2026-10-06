package rerank

import "strings"

func valid(text string) string {
	return strings.ToValidUTF8(text, "�")
}
