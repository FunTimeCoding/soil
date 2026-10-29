package report

import "time"

func New(
	name string,
	size int64,
	modified time.Time,
) *Report {
	return &Report{Name: name, Size: size, Modified: modified}
}
