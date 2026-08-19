package publish

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func New(
	f face.Forge,
	project int64,
	branch string,
	secretAuthority string,
	secretPath string,
) *Publisher {
	return &Publisher{
		forge:           f,
		project:         project,
		branch:          branch,
		secretAuthority: secretAuthority,
		secretPath:      secretPath,
	}
}
