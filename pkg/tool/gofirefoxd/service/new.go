package service

import "github.com/funtimecoding/soil/pkg/firefox/face"

func New(
	c face.FirefoxClient,
	downloadDirectory string,
) *Service {
	return &Service{client: c, downloadDirectory: downloadDirectory}
}
