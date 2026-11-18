package nut

import "strings"

func unquote(v string) string {
	v = strings.TrimSuffix(strings.TrimPrefix(v, `"`), `"`)

	return strings.NewReplacer(`\"`, `"`, `\\`, `\`).Replace(v)
}
