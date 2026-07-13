package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-onnx/ort"
	"github.com/funtimecoding/soil/pkg/errors"
)

func newSession(
	modelPath string,
	sequenceLength int,
) (*rerankSession, error) {
	inputIDs := make([]int64, sequenceLength)
	attentionMask := make([]int64, sequenceLength)
	inputShape := ort.Shape{1, int64(sequenceLength)}
	inputIDsTensor, e := ort.NewTensor[int64](inputShape, inputIDs)

	if e != nil {
		return nil, fmt.Errorf("create input_ids tensor: %w", e)
	}

	attentionMaskTensor, f := ort.NewTensor[int64](inputShape, attentionMask)

	if f != nil {
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create attention_mask tensor: %w", f)
	}

	outputTensor, g := ort.NewEmptyTensor[float32](ort.Shape{1, 1})

	if g != nil {
		errors.PanicOnError(attentionMaskTensor.Destroy())
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create output tensor: %w", g)
	}

	session, h := ort.NewAdvancedSession(
		modelPath,
		[]string{"input_ids", "attention_mask"},
		[]string{"logits"},
		[]ort.Value{inputIDsTensor, attentionMaskTensor},
		[]ort.Value{outputTensor},
		nil,
	)

	if h != nil {
		errors.PanicOnError(outputTensor.Destroy())
		errors.PanicOnError(attentionMaskTensor.Destroy())
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create rerank session: %w", h)
	}

	return &rerankSession{
		inputIDs:            inputIDs,
		attentionMask:       attentionMask,
		inputIDsTensor:      inputIDsTensor,
		attentionMaskTensor: attentionMaskTensor,
		outputTensor:        outputTensor,
		session:             session,
	}, nil
}
