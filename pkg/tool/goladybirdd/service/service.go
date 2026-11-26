package service

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
	"sync"
)

type Service struct {
	client            face.LadybirdClient
	guard             *peer.Guard
	navigation        *navigation.Watcher
	downloadDirectory string
	references        map[string]map[string]string
	mutex             sync.Mutex
}
