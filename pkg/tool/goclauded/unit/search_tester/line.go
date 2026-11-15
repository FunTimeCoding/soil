package search_tester

import "fmt"

func Line(
	identifier string,
	role string,
	timestamp string,
	meta bool,
	content string,
) string {
	return fmt.Sprintf(
		`{"uuid":"%s","type":"%s","timestamp":"%s","isMeta":%t,"message":{"role":"%s","content":%s}}%s`,
		identifier,
		role,
		timestamp,
		meta,
		role,
		content,
		"\n",
	)
}
