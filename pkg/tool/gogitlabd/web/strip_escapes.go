package web

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"

func stripEscapes(s string) string {
	return constant.EscapePattern.ReplaceAllString(s, "")
}
