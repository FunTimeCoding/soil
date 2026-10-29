package report

func New(
	path string,
	tidCount int,
) *Report {
	return &Report{Path: path, TidCount: tidCount}
}
