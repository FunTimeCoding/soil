package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func (s *Store) SwapLabel(
	sessionIdentifier string,
	key string,
	expect string,
	value string,
) (bool, error) {
	match := s.database.Where(
		"session_identifier = ? AND key = ? AND value = ?",
		sessionIdentifier,
		key,
		expect,
	)

	switch {
	case expect == "" && value == "":
		return true, nil
	case expect == "":
		result := s.database.Exec(
			`INSERT INTO label (session_identifier, key, value)
			SELECT ?, ?, ? WHERE NOT EXISTS (
				SELECT 1 FROM label WHERE session_identifier = ? AND key = ?
			)`,
			sessionIdentifier,
			key,
			value,
			sessionIdentifier,
			key,
		)

		return result.RowsAffected > 0, result.Error
	case value == "":
		result := match.Delete(label.Stub())

		return result.RowsAffected > 0, result.Error
	default:
		result := match.Model(label.Stub()).Update(constant.Value, value)

		return result.RowsAffected > 0, result.Error
	}
}
