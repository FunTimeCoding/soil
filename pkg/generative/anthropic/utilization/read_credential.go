package utilization

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"

func ReadCredential() *credential.Credential {
	if !Supported() {
		return nil
	}

	return ParseCredential(rawCredential())
}
