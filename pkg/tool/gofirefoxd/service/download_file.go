package service

import (
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/firefox/download"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) DownloadFile(
	l string,
	path string,
) (*download.Download, error) {
	if l == "" {
		return nil, validation.New(browserConstant.FirefoxLocatorRequired)
	}

	return s.client.DownloadFile(l, path)
}
