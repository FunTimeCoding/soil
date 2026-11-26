package service

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
)

func New(
	c face.LadybirdClient,
	g *peer.Guard,
	w *navigation.Watcher,
	downloadDirectory string,
) *Service {
	return &Service{
		client:            c,
		guard:             g,
		navigation:        w,
		downloadDirectory: downloadDirectory,
		references:        make(map[string]map[string]string),
	}
}
