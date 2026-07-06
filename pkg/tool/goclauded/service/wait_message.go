package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func (s *Service) WaitMessage(
	name string,
	timeout time.Duration,
) ([]message.Message, error) {
	return s.store.WaitMessage(name, timeout)
}
