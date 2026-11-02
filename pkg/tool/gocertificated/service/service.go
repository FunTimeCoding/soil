package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

type Service struct {
	store     *store.Store
	guard     *peer.Guard
	publisher *publish.Publisher
}
