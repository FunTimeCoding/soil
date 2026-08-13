package project

func ReplaceGoImageVersion(
	content string,
	v string,
) string {
	return ReplaceVersionByPrefix(content, v, `image: ([a-z\.\/]+)?golang:`)
}
