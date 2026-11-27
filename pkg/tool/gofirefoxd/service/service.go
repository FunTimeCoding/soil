package service

import "github.com/funtimecoding/soil/pkg/browser/face"

type Service struct {
	client            face.FirefoxClient
	downloadDirectory string
}
