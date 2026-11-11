package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Store) CruiseSettings(
	sessionIdentifiers []string,
) (map[string]*cruise_setting.Setting, error) {
	result := map[string]*cruise_setting.Setting{}

	if len(sessionIdentifiers) == 0 {
		return result, nil
	}

	var rows []cruise_setting.Setting
	e := s.database.Where("session_identifier IN ?", sessionIdentifiers).Find(
		&rows,
	).Error

	if e != nil {
		return nil, e
	}

	for i := range rows {
		result[rows[i].SessionIdentifier] = &rows[i]
	}

	for _, v := range sessionIdentifiers {
		if result[v] == nil {
			result[v] = cruise_setting.New(v, nil)
		}
	}

	return result, nil
}
