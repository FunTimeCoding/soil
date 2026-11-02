package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

func New(
	s *store.Store,
	p face.Client,
	parserGuard *peer.Guard,
	logCachePath string,
	elitePath string,
	outputPath string,
	l *logger.Logger,
) *Service {
	return &Service{
		store:        s,
		parser:       p,
		parserGuard:  parserGuard,
		logger:       l,
		logCachePath: logCachePath,
		elitePath:    elitePath,
		outputPath:   outputPath,
	}
}
