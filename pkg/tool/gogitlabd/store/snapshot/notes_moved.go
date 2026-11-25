package snapshot

func (s *Snapshot) NotesMoved(before *Snapshot) bool {
	return s.Notes != before.Notes || !s.Update.Equal(before.Update)
}
