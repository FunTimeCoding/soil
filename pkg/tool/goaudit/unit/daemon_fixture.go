package unit

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func daemonFixture(run string) *virtual_file_system.System {
	result := virtual_file_system.New()
	result.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	result.WriteString("pkg/tool/gotestd/run.go", run)
	result.WriteString("pkg/tool/gotestd/store/store.go", "package store\n")
	result.WriteString(
		"pkg/tool/gotestd/daemon/daemon.go",
		"package daemon\n\ntype Daemon struct{}\n",
	)
	result.WriteString(
		"pkg/tool/gotestd/service/service.go",
		"package service\n\ntype Service struct{}\n",
	)

	return result
}
