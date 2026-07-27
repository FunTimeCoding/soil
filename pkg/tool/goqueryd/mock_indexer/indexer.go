package mock_indexer

type Indexer struct {
	Pushed    []PushCall
	Deleted   []DeleteCall
	Documents map[string]string
}
