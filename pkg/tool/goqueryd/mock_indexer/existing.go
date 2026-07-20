package mock_indexer

func (i *Indexer) Existing() map[string]string {
	return i.Documents
}
