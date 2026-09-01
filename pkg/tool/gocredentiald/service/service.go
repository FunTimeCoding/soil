package service

import (
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"sync"
	"time"
)

type Service struct {
	path     string
	password string
	client   *keepass.Client
	mutex    sync.Mutex
	backedUp bool
	revealed map[string]bool
	clock    func() time.Time
	logger   *logger.Logger
}
