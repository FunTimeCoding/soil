package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"sync"
)

type Service struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	notifier  *notifier.Notifier
	mutex     sync.RWMutex
	entries   map[string][]*board_entry.Entry
	requests  map[string][]*request_entry.Entry
}
