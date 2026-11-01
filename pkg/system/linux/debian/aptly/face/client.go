package face

type Client interface {
	AddToRepository(
		repository string,
		directory string,
		force bool,
	) error
	Close()
	LatestVersion(
		repository string,
		name string,
	) (string, error)
	Packages(repository string) ([]string, error)
	PublishRepository(
		repositoryName string,
		distribution string,
		architectures []string,
		passphraseFile string,
	) error
	Upload(
		directory string,
		filePath string,
	) error
	UpdatePublish(
		distribution string,
		passphraseFile string,
	) error
	Versions(
		repository string,
		name string,
	) ([]string, error)
}
