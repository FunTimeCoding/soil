package rerank

import "github.com/funtimecoding/soil/pkg/errors"

func destroySession(s *rerankSession) {
	if s.session != nil {
		errors.PanicOnError(s.session.Destroy())
	}

	if s.outputTensor != nil {
		errors.PanicOnError(s.outputTensor.Destroy())
	}

	if s.attentionMaskTensor != nil {
		errors.PanicOnError(s.attentionMaskTensor.Destroy())
	}

	if s.inputIDsTensor != nil {
		errors.PanicOnError(s.inputIDsTensor.Destroy())
	}
}
