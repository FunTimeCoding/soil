package project

import (
	"fmt"
	"regexp"
)

func ReplaceVersionByPrefix(
	content string,
	v string,
	prefix string,
) string {
	return regexp.MustCompile(
		fmt.Sprintf(`(%s)([0-9]+\.[0-9]+\.[0-9]+)`, prefix),
	).ReplaceAllString(content, fmt.Sprintf("${1}%s", v))
}
