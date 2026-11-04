package store

func (s *Store) InsertContent(
	hash string,
	content string,
	now string,
) error {
	_, e := s.database.Exec(
		"INSERT OR IGNORE INTO content (hash, body, created_at) VALUES (?, ?, ?)",
		hash,
		content,
		now,
	)

	return e
}
