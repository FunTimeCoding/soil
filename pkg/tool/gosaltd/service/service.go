package service

import (
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
	store    *store.Store
	reaper   *reaper.Reaper
	runner   *runner.Runner
}
