package store_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func TagFilter(value string) map[string]string {
	return map[string]string{constant.FixtureTagKey: value}
}
