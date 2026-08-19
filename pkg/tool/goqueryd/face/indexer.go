package face

type Indexer interface {
	Push(
		collection string,
		path string,
		body string,
		metadata map[string][]string,
	) error
	MustPush(
		collection string,
		path string,
		body string,
		metadata map[string][]string,
	)
	Existing(collection string) map[string]string
	Delete(
		collection string,
		path string,
	) error
	MustDelete(
		collection string,
		path string,
	)
}
