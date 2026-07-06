package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Store) UpdateCompletionSequence(
	identifier uint,
	sequence int,
) {
	s.database.Model(completion.Stub()).Where(
		"identifier = ?",
		identifier,
	).Update("sequence", sequence)
}
