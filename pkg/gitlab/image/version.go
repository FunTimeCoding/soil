package image

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"log"
	"strings"
)

func (i *Image) Version() string {
	result := strings.Split(i.Path, ":")[1]

	if result == "" {
		log.Panicf("empty version: %+v", result)
	}

	if !strings.HasPrefix(result, constant.VersionPrefix) {
		return key_value.Empty(constant.VersionPrefix, result)
	}

	return result
}
