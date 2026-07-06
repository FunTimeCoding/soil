package flag

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"

type Flag struct {
	Identifier string
	By         []*client.Client
}
