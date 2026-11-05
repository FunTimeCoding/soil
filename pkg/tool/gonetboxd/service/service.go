package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

type Service struct {
	client face.NetBoxClient
	guard  *peer.Guard
	store  *store.Store
}
