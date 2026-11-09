package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"time"
)

type Service struct {
	store     *store.Store
	guard     *peer.Guard
	publisher *publish.Publisher
	clock     func() time.Time
}
