package service

import (
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func (s *Service) EnableRedirection(instance string) ([]string, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	r, f := c.Redirection()

	if f != nil {
		return nil, f
	}

	var result []string

	if !r.ListenerEnabled {
		if g := c.EnableListener(); g != nil {
			return result, g
		}

		result = append(result, constant.ListenerWrite)
	}

	v, h := c.ScreenService()

	if h != nil {
		return result, h
	}

	if v.EnabledState != band.ScreenServiceEnabled &&
		v.EnabledState != band.ScreenServiceOffline {
		if i := c.EnableScreenService(); i != nil {
			return result, i
		}

		result = append(result, constant.ScreenServiceWrite)
	}

	return result, nil
}
