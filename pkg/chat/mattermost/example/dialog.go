package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Dialog() {
	w := http.NewServeMux()
	w.HandleFunc(webConstant.DialogPath, handleDialog)
	s := web.ListenAsynchronous(w)
	m := common.Mattermost()
	defer m.Close()

	if true {
		dialogPostButtons(m)
	}

	if false {
		dialogElements()
	}

	if false {
		dialogOpen(m)
	}

	if false {
		dialogSubmit(m)
	}

	console.Line("Cancel server with Ctrl-C")
	system.KillSignalBlock()
	console.Line("Shutdown server")
	web.GracefulShutdown(context.Background(), s, false)
}
