package service

import "github.com/funtimecoding/soil/pkg/tool/gosentryd/face"

type Service struct {
	client       face.SentrySource
	organization string
}
