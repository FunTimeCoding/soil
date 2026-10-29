package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	database      *gorm.DB
	store         *store.Store
	parser        face.Client
	parserGuard   *peer.Guard
	authorization *client.Client
	logger        *logger.Logger
	logCachePath  string
	elitePath     string
	outputPath    string
	closeOnce     sync.Once
}
