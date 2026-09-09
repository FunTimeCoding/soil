package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

type Server struct {
	client   netbox.NetboxSource
	store    *store.Store
	reporter face.Reporter
}
