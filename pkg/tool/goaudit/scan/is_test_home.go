package scan

import "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"

func isTestHome(name string) bool {
	return name == constant.UnitDirectory ||
		name == constant.IntegrationDirectory
}
