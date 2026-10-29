package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
	store    *store.Store
	reporter face.Reporter
}
