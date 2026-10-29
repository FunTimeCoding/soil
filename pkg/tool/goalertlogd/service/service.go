package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Service struct {
	database     *gorm.DB
	store        *store.Store
	alertmanager face.Client
	guard        *peer.Guard
	retention    time.Duration
	closeOnce    sync.Once
}
