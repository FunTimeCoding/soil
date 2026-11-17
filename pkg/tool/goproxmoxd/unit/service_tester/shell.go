package service_tester

import (
	"github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func (o *Tester) shell(i *instance.Instance) (face.SSHClient, error) {
	if !o.redial {
		return o.Instance(i.Name).Shell, nil
	}

	result := mock_client.New(o.t)
	o.dialled = append(o.dialled, result)

	return result, nil
}
