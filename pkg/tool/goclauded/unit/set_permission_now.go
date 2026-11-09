package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"

func setPermissionNow(x *service_tester.Tester) {
	x.Store.Advance(permissionNow().Sub(x.Store.Clock()()))
}
