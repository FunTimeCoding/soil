package service

import (
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.HabiticaClient
	guard  *peer.Guard
}
