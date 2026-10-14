package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
)

func (s *Store) CruiseSetting(
	sessionIdentifier string,
) (*cruise_setting.Setting, error) {
	var result cruise_setting.Setting
	found := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&result)

	if found.Error != nil {
		return nil, found.Error
	}

	if found.RowsAffected == 0 {
		return cruise_setting.New(sessionIdentifier, constant.CruiseOff, 0), nil
	}

	return &result, nil
}
