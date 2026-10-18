package mock_indexer

func (i *Indexer) Existing(collection string) (map[string]string, error) {
	return i.Documents, nil
}
