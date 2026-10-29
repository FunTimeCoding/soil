package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
)

type Service struct {
	database      *gorm.DB
	store         *store.Store
	forge         face.Forge
	guard         *peer.Guard
	publisher     *publish.Publisher
	authorization *client.Client
}
