package service

import "github.com/funtimecoding/soil/pkg/browser/firefox/download"

func (s *Service) DownloadStatus(identifier int) ([]*download.Download, error) {
	return s.client.DownloadStatus(identifier)
}
