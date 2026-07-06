package flag

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"

func (f *Flag) RemoveClient(v *client.Client) {
	for i, c := range f.By {
		if c == v {
			f.By = append(f.By[:i], f.By[i+1:]...)

			break
		}
	}
}
