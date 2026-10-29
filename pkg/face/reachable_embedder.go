package face

type ReachableEmbedder interface {
	Embedder
	Reachable
	Closer
}
