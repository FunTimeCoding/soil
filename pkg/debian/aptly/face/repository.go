package face

type Repository interface {
	Versions(
		repository string,
		name string,
	) ([]string, error)
	LatestVersion(
		repository string,
		name string,
	) (string, error)
}
