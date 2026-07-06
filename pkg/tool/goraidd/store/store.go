package store

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"gorm.io/gorm"
)

type Store struct {
	mapper       *gorm.DB
	logger       *logger.Logger
	recovery     *recovery.Recovery
	logCachePath string
	elitePath    string
	stop         chan struct{}
}
