package label

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
)

type Service struct {
	store    *store.Store
	notifier face.Notifier
}
