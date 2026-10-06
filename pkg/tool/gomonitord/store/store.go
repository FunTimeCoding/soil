package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"gorm.io/gorm"
)

type Store struct {
	database *gorm.DB
	notifier face.EventNotifier
}
