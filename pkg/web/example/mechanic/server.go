package mechanic

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/view"
	"sync"
)

type Server struct {
	view     *view.View
	notifier face.EventNotifier
	mutex    sync.Mutex
	count    int
	pulse    int
}
