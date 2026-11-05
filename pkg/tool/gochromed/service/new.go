package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

func New(
	c face.ChromiumSource,
	g *peer.Guard,
	downloadDirectory string,
) *Service {
	return &Service{
		source:            c,
		guard:             g,
		downloadDirectory: downloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
	}
}
