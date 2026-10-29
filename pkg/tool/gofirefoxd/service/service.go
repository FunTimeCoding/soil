package service

import "github.com/funtimecoding/soil/pkg/firefox/face"

type Service struct {
	client            face.Client
	downloadDirectory string
}
