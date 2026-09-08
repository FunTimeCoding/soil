package utilization

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/keychain"
)

func keychainRaw() string {
	return keychain.Password(constant.AnthropicCredentialService)
}
