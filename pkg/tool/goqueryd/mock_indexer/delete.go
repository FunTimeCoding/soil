package mock_indexer

func (i *Indexer) Delete(collection string, path string) error {
	i.Deleted = append(
		i.Deleted,
		DeleteCall{Collection: collection, Path: path},
	)

	return nil
}
