package service

import "github.com/funtimecoding/soil/pkg/directory"

func New(d *directory.Client) *Service {
	return &Service{directory: d}
}
