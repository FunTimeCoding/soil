package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"time"
)

func (s *Service) Client(instance string) (face.ProxmoxClient, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.clients[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	var o []proxmox.Option

	if i.User != "" && i.Password != "" {
		o = append(o, proxmox.WithUser(i.User, i.Password))
	} else if i.Token != "" && i.Secret != "" {
		o = append(o, proxmox.WithToken(i.Token, i.Secret))
	}

	if i.Port > 0 {
		o = append(o, proxmox.WithPort(i.Port))
	}

	if i.Insecure {
		o = append(o, proxmox.WithInsecure())
	}

	if i.Timeout != "" {
		if d, e := time.ParseDuration(i.Timeout); e == nil {
			o = append(o, proxmox.WithTimeout(d))
		}
	}

	c := proxmox.New(i.Host, o...)
	s.clients[instance] = c

	return c, nil
}
