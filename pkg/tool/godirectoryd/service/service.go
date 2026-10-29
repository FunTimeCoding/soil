package service

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
)

type Service struct {
	directory     face.Client
	guard         *peer.Guard
	authorization *client.Client
}
