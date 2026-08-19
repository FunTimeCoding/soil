package mock_indexer

type PushCall struct {
	Collection string
	Name       string
	Body       string
	Metadata   map[string][]string
}
