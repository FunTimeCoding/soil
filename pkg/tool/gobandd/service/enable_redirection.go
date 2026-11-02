package service

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
)

func (s *Service) EnableRedirection(instance string) ([]string, error) {
	return guarded(
		s,
		instance,
		func(c face.BandClient) ([]string, error) {
			r, e := c.Redirection()

			if e != nil {
				return nil, e
			}

			var result []string

			if !r.ListenerEnabled {
				if f := c.EnableListener(); f != nil {
					return result, f
				}

				result = append(result, constant.ListenerWrite)
			}

			v, g := c.ScreenService()

			if g != nil {
				return result, g
			}

			if v.EnabledState != bandConstant.ScreenServiceEnabled &&
				v.EnabledState != bandConstant.ScreenServiceOffline {
				if h := c.EnableScreenService(); h != nil {
					return result, h
				}

				result = append(result, constant.ScreenServiceWrite)
			}

			return result, nil
		},
	)
}
