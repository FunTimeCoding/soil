package message

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store    *store.Store
	notifier face.Notifier
	stop     *stopper.Stopper
	clock    func() time.Time
	reporter soilFace.Reporter
}
