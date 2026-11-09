package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"time"
)

type Service struct {
	store    *store.Store
	clock    func() time.Time
	reporter face.Reporter
}
