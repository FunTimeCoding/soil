package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

type Server struct {
	client   *netbox.Client
	store    *store.Store
	reporter face.Reporter
}
