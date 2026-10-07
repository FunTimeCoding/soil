package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/raid/model/player_fight_statistic"
	"github.com/funtimecoding/soil/pkg/raid/model/raid"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	logCachePath string,
	elitePath string,
	l *logger.Logger,
	r face.Reporter,
) *Store {
	errors.PanicOnError(m.AutoMigrate(raid.New()))
	errors.PanicOnError(m.AutoMigrate(fight.New()))
	errors.PanicOnError(m.AutoMigrate(player_fight_statistic.New()))
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
