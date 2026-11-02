package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"gorm.io/gorm"
)

type Tester struct {
	Service  *service.Service
	Notifier *notifier.Notifier
	database *gorm.DB
}
