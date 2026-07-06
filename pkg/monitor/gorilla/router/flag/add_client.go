package flag

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"

func (f *Flag) AddClient(v *client.Client) {
	f.By = append(f.By, v)
}
