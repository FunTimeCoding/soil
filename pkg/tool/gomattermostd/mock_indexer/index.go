package mock_indexer

func (i *Indexer) Index(root string) {
	i.indexed = append(i.indexed, root)
}
