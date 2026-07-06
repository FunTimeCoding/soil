package gorenovate

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func parseConfiguration() *Configuration {
	if !system.FileExists(RenovateFile) {
		return nil
	}

	var result Configuration
	notation.MustDecodeBytes(
		system.ReadBytesUnsafe(RenovateFile),
		&result,
		false,
	)

	return &result
}
