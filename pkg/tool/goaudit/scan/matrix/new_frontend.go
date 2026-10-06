package matrix

func NewFrontend(
	name string,
	repo string,
	theme string,
	favicon bool,
) *Frontend {
	return &Frontend{Name: name, Repo: repo, Theme: theme, Favicon: favicon}
}
