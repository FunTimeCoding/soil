package release

func New(
	codename string,
	major int64,
	minor int64,
) *Release {
	return &Release{Codename: codename, Major: major, Minor: minor}
}
