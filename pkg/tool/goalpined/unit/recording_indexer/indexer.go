package recording_indexer

type Indexer struct {
	Directories []string
	KeyPaths    []string
	Failure     error
}
