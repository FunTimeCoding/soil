package face

type Indexer interface {
	Push(
		path string,
		body string,
		metadata map[string]string,
	) error
	MustPush(
		path string,
		body string,
		metadata map[string]string,
	)
	Existing() map[string]string
	Delete(path string) error
	MustDelete(path string)
}
