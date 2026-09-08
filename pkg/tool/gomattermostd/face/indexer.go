package face

type Indexer interface {
	Index(root string)
	Forget(root string)
}
