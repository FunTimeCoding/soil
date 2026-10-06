package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"

func (s *Store) findActiveDocument(
	collection string,
	path string,
) *record.ActiveDocument {
	row := s.database.QueryRow(
		"SELECT identifier, hash, title FROM document WHERE collection = ? AND path = ? AND active = 1",
		collection,
		path,
	)
	var d record.ActiveDocument
	e := row.Scan(&d.Identifier, &d.Hash, &d.Title)

	if e != nil {
		return nil
	}

	return &d
}
