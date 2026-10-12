package service

import "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"

func New(
	c face.SentrySource,
	organization string,
) *Service {
	return &Service{client: c, organization: organization}
}
