package store_recorder

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Recorder struct {
	store    *store.Store
	reporter face.Reporter
}
