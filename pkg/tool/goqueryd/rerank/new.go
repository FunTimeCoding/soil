package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-onnx/ort"
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func New(
	modelPath string,
	tokenizerPath string,
) (*Reranker, error) {
	if !ort.IsInitialized() {
		if e := ort.InitializeEnvironmentWithBootstrap(); e != nil {
			return nil, fmt.Errorf("onnx runtime init: %w", e)
		}
	}

	tokenizer, e := tokenizers.FromFile(
		tokenizerPath,
		tokenizers.WithTruncation(
			uintptr(constant.DefaultSequenceLength),
			tokenizers.TruncationDirectionRight,
			tokenizers.TruncationStrategyLongestFirst,
		),
		tokenizers.WithPadding(
			true,
			tokenizers.PaddingStrategy{
				Tag:       tokenizers.PaddingStrategyFixed,
				FixedSize: uintptr(constant.DefaultSequenceLength),
			},
		),
	)

	if e != nil {
		return nil, fmt.Errorf("load tokenizer: %w", e)
	}

	return &Reranker{
		modelPath:      modelPath,
		sequenceLength: constant.DefaultSequenceLength,
		tokenizer:      tokenizer,
		sessions:       make(map[int]*rerankSession),
	}, nil
}
