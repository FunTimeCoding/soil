package identity

import "github.com/funtimecoding/soil/pkg/stamp"

func (t *Tool) WithStamp(
	version string,
	gitHash string,
	buildDate string,
) *Tool {
	t.stamp = stamp.New(version, gitHash, buildDate)

	return t
}
