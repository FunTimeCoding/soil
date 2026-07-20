package mock_indexer

type Indexer struct {
	Pushed    []PushCall
	Deleted   []string
	Documents map[string]string
}
