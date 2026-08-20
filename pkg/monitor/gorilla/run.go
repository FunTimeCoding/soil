package gorilla

import (
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log"
	"net/http"
)

func Run(address string) {
	r := router.New()
	m := http.NewServeMux()
	m.HandleFunc(constant.MonitorPath, r.Monitor)
	m.HandleFunc(constant.EchoPath, echo)
	m.HandleFunc(constant.RootPath, home)
	log.Printf("listen on %s\n", address)
	web.ListenAddress(m, address)
}
