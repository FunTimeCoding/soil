package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"gorm.io/gorm"
)

type Service struct {
	database   *gorm.DB
	store      *store.Store
	reaper     *reaper.Reaper
	downstream []face.Client
	guards     []*peer.Guard
	runner     *runner.Runner
}
