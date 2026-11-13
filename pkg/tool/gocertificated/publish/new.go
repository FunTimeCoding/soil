package publish

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func New(
	f face.GitLabClient,
	projectPath string,
	branch string,
	authorityDirectory string,
	secretAuthority string,
	secretPath string,
) *Publisher {
	return &Publisher{
		forge:              f,
		projectPath:        projectPath,
		branch:             branch,
		authorityDirectory: authorityDirectory,
		secretAuthority:    secretAuthority,
		secretPath:         secretPath,
	}
}
