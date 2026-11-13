package tree_entry

func (e *Entry) Same(other *Entry) bool {
	return other != nil &&
		e.Size == other.Size &&
		e.Modified.Equal(other.Modified) &&
		e.Directory == other.Directory &&
		e.Linked == other.Linked
}
