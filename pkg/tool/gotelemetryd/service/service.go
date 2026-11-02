package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Service struct {
	store    *store.Store
	reporter face.Reporter
}
