package store

func (s *Store) Summaries() ([]Summary, error) {
	var result []Summary

	return result, s.mapper.
		Model(NewClick()).
		Select("label, count(*) as count, max(created_at) as last").
		Group(LabelColumn).
		Order("count DESC").
		Scan(&result).Error
}
