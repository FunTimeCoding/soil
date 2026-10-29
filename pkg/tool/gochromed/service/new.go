package service

import "github.com/funtimecoding/soil/pkg/tool/gochromed/option"

func New(o *option.Chrome) *Service {
	return &Service{
		source:            o.Source(),
		downloadDirectory: o.DownloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
	}
}
