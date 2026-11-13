package publish

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"sync"
)

type Publisher struct {
	forge              face.GitLabClient
	projectPath        string
	project            int64
	mutex              sync.Mutex
	branch             string
	authorityDirectory string
	secretAuthority    string
	secretPath         string
}
