package service

import "github.com/funtimecoding/soil/pkg/browser/face"

func New(
	c face.FirefoxClient,
	downloadDirectory string,
) *Service {
	return &Service{client: c, downloadDirectory: downloadDirectory}
}
