package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"
	"gorm.io/gorm"
	"time"
)

func (s *Store) GetLoginSession(
	identifier string,
) (*login_session.LoginSession, error) {
	var row login_session.LoginSession
	e := s.mapper.Where("identifier = ?", identifier).First(&row).Error

	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, e
	}

	if time.Since(row.CreatedAt) > constant.LoginSessionTimeToLive {
		errors.PanicOnError(s.DeleteLoginSession(identifier))

		return nil, nil
	}

	return &row, nil
}
