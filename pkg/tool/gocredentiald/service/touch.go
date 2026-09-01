package service

import (
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
)

func (s *Service) touch(entry *gokeepasslib.Entry) {
	at := wrappers.Now()
	at.Time = s.clock()
	entry.Times.LastModificationTime = &at
}
