package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Service struct {
	mutex     sync.Mutex
	database  *gorm.DB
	store     *store.Store
	logger    *logger.Logger
	predicate string
	retention time.Duration
	clock     func() time.Time
	command   func(s ...string) (string, error)
	stream    func(s ...string) face.Stream
	last      map[string]string
}
