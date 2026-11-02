package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

type Service struct {
	store        *store.Store
	parser       face.Client
	parserGuard  *peer.Guard
	logger       *logger.Logger
	logCachePath string
	elitePath    string
	outputPath   string
}
