package named_entry

func (e *Entry) IsDir() bool {
	return e.entry.Directory
}
