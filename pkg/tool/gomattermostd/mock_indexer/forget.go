package mock_indexer

func (i *Indexer) Forget(root string) {
	i.forgotten = append(i.forgotten, root)
}
