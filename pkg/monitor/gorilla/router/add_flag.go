package router

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/flag"

func (r *Router) AddFlag(f *flag.Flag) {
	r.Flag = append(r.Flag, f)
}
