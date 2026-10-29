package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
	client   face.Client
	store    *store.Store
}
