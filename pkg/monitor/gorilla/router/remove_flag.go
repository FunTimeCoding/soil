package router

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/flag"

func (r *Router) RemoveFlag(f *flag.Flag) {
	for i, l := range r.Flag {
		if l == f {
			r.Flag = append(r.Flag[:i], r.Flag[i+1:]...)

			break
		}
	}
}
