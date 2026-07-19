package service

import (
	"fmt"
	"strings"
)

func ReplaceUnique(
	content string,
	oldText string,
	newText string,
) (string, error) {
	count := strings.Count(content, oldText)

	if count == 0 {
		return "", fmt.Errorf(
			"old_text not found in page. Re-read the page with confluence_get_page and try again",
		)
	}

	if count > 1 {
		return "", fmt.Errorf(
			"old_text found %d times, must be unique. Provide more surrounding context to match exactly once",
			count,
		)
	}

	return strings.Replace(content, oldText, newText, 1), nil
}
