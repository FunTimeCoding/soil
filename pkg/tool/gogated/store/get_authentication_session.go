package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"gorm.io/gorm"
	"time"
)

func (s *Store) GetAuthenticationSession(
	identifier string,
) (*authentication_session.AuthenticationSession, error) {
	var row authentication_session.AuthenticationSession
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, e
	}

	now := time.Now()
	idle := now.Sub(row.LastUsedAt) > constant.AuthenticationIdleTimeToLive
	absolute := now.Sub(row.CreatedAt) >
		constant.AuthenticationAbsoluteTimeToLive

	if idle || absolute {
		errors.PanicOnError(s.DeleteAuthenticationSession(identifier))

		return nil, nil
	}

	row.LastUsedAt = now
	e = s.mapper.Model(authentication_session.Stub()).
		Where("identifier = ?", identifier).
		Update("last_used_at", now).Error

	if e != nil {
		return nil, e
	}

	return &row, nil
}
