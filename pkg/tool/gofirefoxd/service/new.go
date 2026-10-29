package service

import "github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"

func New(o *option.Firefox) *Service {
	return &Service{
		client:            o.Client(),
		downloadDirectory: o.DownloadDirectory,
	}
}
