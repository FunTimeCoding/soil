package tokenizer

import "strings"

func Valid(text string) string {
	return strings.ToValidUTF8(text, "�")
}
