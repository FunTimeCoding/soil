package image

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"log"
	"strings"
)

func (i *Image) Version() string {
	result := strings.Split(i.Path, ":")[1]

	if result == "" {
		log.Panicf("empty version: %+v", result)
	}

	if !strings.HasPrefix(result, gitConstant.VersionPrefix) {
		return key_value.Empty(gitConstant.VersionPrefix, result)
	}

	return result
}
