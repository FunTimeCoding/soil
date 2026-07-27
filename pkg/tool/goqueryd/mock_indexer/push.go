package mock_indexer

func (i *Indexer) Push(
	collection string,
	name string,
	body string,
	metadata map[string]string,
) error {
	i.Pushed = append(
		i.Pushed,
		PushCall{
			Collection: collection,
			Name:       name,
			Body:       body,
			Metadata:   metadata,
		},
	)

	return nil
}
