package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/raid"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	logCachePath string,
	elitePath string,
	l *logger.Logger,
	r face.Reporter,
) *Store {
	errors.PanicOnError(m.AutoMigrate(raid.NewRaid()))
	errors.PanicOnError(m.AutoMigrate(raid.NewFight()))
	errors.PanicOnError(m.AutoMigrate(raid.NewPlayerFightStatistic()))
	s := &Store{
		mapper:       m,
		logger:       l,
		recovery:     recovery.New(l, r),
		logCachePath: logCachePath,
		elitePath:    elitePath,
		stop:         make(chan struct{}),
	}
	s.syncLogCache()

	return s
}
