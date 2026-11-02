package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"time"
)

func New(
	c face.ItermSource,
	keyInterval time.Duration,
) *Service {
	return &Service{source: c, keyInterval: keyInterval}
}
