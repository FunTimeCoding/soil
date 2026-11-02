package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/download"
)

func (s *Service) DownloadFile(
	l string,
	path string,
) (*download.Download, error) {
	if l == "" {
		return nil, validation.New(firefoxConstant.LocatorRequired)
	}

	return s.client.DownloadFile(l, path)
}
