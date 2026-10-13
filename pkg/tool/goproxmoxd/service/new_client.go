package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"time"
)

func newClient(i *instance.Instance) (face.ProxmoxClient, error) {
	var o []proxmox.Option

	if i.User != "" && i.Password != "" {
		o = append(o, proxmox.WithUser(i.User, i.Password))
	} else if i.Token != "" && i.Secret != "" {
		o = append(o, proxmox.WithToken(i.Token, i.Secret))
	}

	if i.Port > 0 {
		o = append(o, proxmox.WithPort(i.Port))
	}

	if i.Untrusted {
		o = append(o, proxmox.WithUntrusted())
	}

	if i.Timeout != "" {
		if d, e := time.ParseDuration(i.Timeout); e == nil {
			o = append(o, proxmox.WithTimeout(d))
		}
	}

	return proxmox.New(i.Host, o...), nil
}
