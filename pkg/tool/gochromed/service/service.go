package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"sync"
)

type Service struct {
	source            face.ChromiumSource
	guard             *peer.Guard
	downloadDirectory string
	snapshotCache     map[string]map[string]int64
	mutex             sync.Mutex
}
