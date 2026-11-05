package service

import (
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.OpnsenseClient
	guard  *peer.Guard
}
