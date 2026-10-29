package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
)

func (s *Service) CreateMark(
	label string,
	note string,
) (*mark.Mark, error) {
	if label == "" {
		return nil, validation.New(constant.LabelRequired)
	}

	result := mark.New(s.clock(), label, note)

	if e := s.store.CreateMark(result); e != nil {
		return nil, e
	}

	return result, nil
}
