package gorenovate

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gorenovate/constant"
)

func parseConfiguration() *Configuration {
	if !system.FileExists(constant.RenovateFile) {
		return nil
	}

	var result Configuration
	notation.MustDecodeBytes(
		system.ReadBytesUnsafe(constant.RenovateFile),
		&result,
		false,
	)

	return &result
}
