package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/download"
)

func (s *Service) DownloadFile(
	l string,
	path string,
) (*download.Download, error) {
	if l == "" {
		return nil, validation.New(constant.LocatorRequired)
	}

	return s.client.DownloadFile(l, path)
}
