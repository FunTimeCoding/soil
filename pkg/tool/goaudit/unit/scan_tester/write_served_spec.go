package scan_tester

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func WriteServedSpec(
	v *virtual_file_system.System,
	spec string,
) {
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString("pkg/tool/gotestd/server/server.go", "package server\n")
	v.WriteString("pkg/tool/gotestd/generated/server/openapi.yaml", spec)
}
