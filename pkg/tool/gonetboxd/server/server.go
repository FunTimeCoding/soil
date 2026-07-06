package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/netbox"
)

type Server struct {
	client   *netbox.Client
	reporter face.Reporter
}
