package service

import (
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"gorm.io/gorm"
)

type Service struct {
	database  *gorm.DB
	store     *store.Store
	reaper    *reaper.Reaper
	inventory face.Client
	runner    *runner.Runner
}
