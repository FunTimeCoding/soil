package publish

import "github.com/funtimecoding/soil/pkg/gitlab/face"

type Publisher struct {
	forge              face.GitLabClient
	project            int64
	branch             string
	authorityDirectory string
	secretAuthority    string
	secretPath         string
}
