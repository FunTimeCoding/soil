package session_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/environment"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func memoryBase() string {
	host := environment.Fallback(gomemorydConstant.HostEnvironment, "")
	port := environment.Fallback(gomemorydConstant.PortEnvironment, "")

	if host == "" || port == "" {
		return ""
	}

	return join.Empty("http://", host, ":", port, "/memories/")
}
