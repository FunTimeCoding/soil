package store_recorder

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Recorder {
	return &Recorder{store: s, reporter: r}
}
