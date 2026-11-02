package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"sync/atomic"
	"time"
)

type Service struct {
	store        *store.Store
	alertmanager prometheus.AlertmanagerClient
	guard        *peer.Guard
	retention    time.Duration
	notifier     face.EventNotifier
	lastPoll     atomic.Value
}
