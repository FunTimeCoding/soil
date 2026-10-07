package rerank_session

import "github.com/amikos-tech/pure-onnx/ort"

type Session struct {
	InputIDs            []int64
	AttentionMask       []int64
	InputIDsTensor      *ort.Tensor[int64]
	AttentionMaskTensor *ort.Tensor[int64]
	OutputTensor        *ort.Tensor[float32]
	Session             *ort.AdvancedSession
}
