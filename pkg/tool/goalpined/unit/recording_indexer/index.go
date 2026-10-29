package recording_indexer

func (i *Indexer) Index(
	directory string,
	keyPath string,
) error {
	i.Directories = append(i.Directories, directory)
	i.KeyPaths = append(i.KeyPaths, keyPath)

	return i.Failure
}
