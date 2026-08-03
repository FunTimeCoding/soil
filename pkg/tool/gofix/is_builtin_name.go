package gofix

import "github.com/funtimecoding/soil/pkg/tool/gofix/constant"

func isBuiltinName(name string) bool {
	return constant.BuiltinNames[name]
}
