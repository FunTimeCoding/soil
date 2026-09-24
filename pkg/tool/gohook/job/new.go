package job

func New(
	run string,
	paths ...string,
) *Job {
	return &Job{Run: run, Paths: paths}
}
