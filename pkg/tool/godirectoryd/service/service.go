package service

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	directory face.DirectoryClient
	guard     *peer.Guard
}
