package service

import "github.com/funtimecoding/soil/pkg/tool/gochromed/face"

func New(
	c face.ChromiumSource,
	downloadDirectory string,
) *Service {
	return &Service{
		source:            c,
		downloadDirectory: downloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
	}
}
