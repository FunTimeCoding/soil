package store

func (s *Store) ActiveDocumentHash(
	collection string,
	path string,
) string {
	d := s.findActiveDocument(collection, path)

	if d == nil {
		return ""
	}

	return d.hash
}
