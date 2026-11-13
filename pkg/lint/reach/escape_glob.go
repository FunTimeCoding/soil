package reach

import "strings"

func escapeGlob(s string) string {
	return strings.NewReplacer(
		`\`,
		`\\`,
		`*`,
		`\*`,
		`?`,
		`\?`,
		`[`,
		`\[`,
	).Replace(s)
}
