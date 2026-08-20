package publish

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func New(
	f face.Forge,
	project int64,
	branch string,
	authorityDirectory string,
	secretAuthority string,
	secretPath string,
) *Publisher {
	return &Publisher{
		forge:              f,
		project:            project,
		branch:             branch,
		authorityDirectory: authorityDirectory,
		secretAuthority:    secretAuthority,
		secretPath:         secretPath,
	}
}
