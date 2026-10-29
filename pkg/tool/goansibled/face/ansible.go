package face

import "github.com/funtimecoding/soil/pkg/provision/face"

type Ansible interface {
	face.Provision
	Playbooks() []string
}
