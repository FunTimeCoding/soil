package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func isHyphenatedToken(term string) bool {
	return constant.HyphenatedPattern.MatchString(term)
}
