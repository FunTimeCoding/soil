package rerank_session

import "github.com/funtimecoding/soil/pkg/errors"

func Destroy(s *Session) {
	if s.Session != nil {
		errors.PanicOnError(s.Session.Destroy())
	}

	if s.OutputTensor != nil {
		errors.PanicOnError(s.OutputTensor.Destroy())
	}

	if s.AttentionMaskTensor != nil {
		errors.PanicOnError(s.AttentionMaskTensor.Destroy())
	}

	if s.InputIDsTensor != nil {
		errors.PanicOnError(s.InputIDsTensor.Destroy())
	}
}
