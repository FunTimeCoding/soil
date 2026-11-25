package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/note_writer"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	watchNotifier "github.com/funtimecoding/soil/pkg/watch/notifier"
	"github.com/funtimecoding/soil/pkg/watch/pending_digest"
	"sync"
)

type Service struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	notifier    *notifier.Notifier
	mutex       sync.RWMutex
	entries     map[string][]*board_entry.Entry
	requests    map[string][]*request_entry.Entry
	writers     *note_writer.Writers
	store       *store.Store
	digests     *watchNotifier.Notifier
	roster      face.Client
	rosterGuard *peer.Guard
	pending     *pending_digest.Pending
	outages     *outage_tracker.Tracker
	logger      *logger.Logger
}
