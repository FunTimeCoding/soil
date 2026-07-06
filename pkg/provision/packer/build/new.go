package build

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"log"
)

func New(
	architecture string,
	characterDevicePort int,
	username string,
	password string,
) *Build {
	switch architecture {
	case constant.AMD64:
		// pass
	default:
		log.Panicf("unsupported architecture: %s", architecture)
	}

	b := &Build{
		architecture:        architecture,
		CharacterDevicePort: characterDevicePort,
		username:            username,
		password:            password,
	}
	b.SetProvisioner([]string{"echo hello"})

	return b
}
