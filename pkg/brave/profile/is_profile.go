package profile

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"os"
	"strings"
)

func IsProfile(f os.DirEntry) bool {
	if !f.IsDir() {
		return false
	}

	if f.Name() == constant.DefaultProfile {
		return true
	}

	if strings.HasPrefix(f.Name(), constant.ProfilePrefix) {
		return true
	}

	return false
}
