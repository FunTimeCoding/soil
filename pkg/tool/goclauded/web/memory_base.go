package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func memoryBase() string {
	host := environment.Fallback(constant.HostEnvironment, "")
	port := environment.Fallback(constant.PortEnvironment, "")

	if host == "" || port == "" {
		return ""
	}

	return join.Empty("http://", host, ":", port, "/memories/")
}
