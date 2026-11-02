package service

import "github.com/funtimecoding/soil/pkg/tool/gogithubd/face"

func New(
	s face.RunSource,
	owner string,
) *Service {
	return &Service{source: s, owner: owner}
}
