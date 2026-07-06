package router

import (
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/flag"
)

type Router struct {
	Client []*client.Client
	Flag   []*flag.Flag
}
