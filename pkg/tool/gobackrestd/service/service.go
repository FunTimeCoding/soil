package service

import (
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.BackrestClient
	guard  *peer.Guard
}
