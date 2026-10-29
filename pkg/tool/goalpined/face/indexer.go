package face

type Indexer interface {
	Index(
		directory string,
		keyPath string,
	) error
}
