package service_tester

import (
	"github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func (s *Shells) build(*instance.Instance) (face.SSHClient, error) {
	if s.given != nil {
		return s.given, nil
	}

	result := mock_client.New()
	s.Built = append(s.Built, result)

	return result, nil
}
