package unit

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func tagFilter(value string) map[string]string {
	return map[string]string{constant.FixtureTagKey: value}
}
