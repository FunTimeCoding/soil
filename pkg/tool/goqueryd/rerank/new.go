package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-onnx/ort"
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func New(
	name string,
	directory string,
	sequenceLength int,
) (*Reranker, error) {
	if !ort.IsInitialized() {
		if e := ort.InitializeEnvironmentWithBootstrap(); e != nil {
			return nil, fmt.Errorf("onnx runtime init: %w", e)
		}
	}

	path := join.Join(directory, constant.RerankTokenizerFile)
	tokenizer, e := tokenizers.FromFile(
		path,
		tokenizers.WithTruncation(
			uintptr(sequenceLength),
			tokenizers.TruncationDirectionRight,
			tokenizers.TruncationStrategyLongestFirst,
		),
		tokenizers.WithPadding(
			true,
			tokenizers.PaddingStrategy{
				Tag:       tokenizers.PaddingStrategyFixed,
				FixedSize: uintptr(sequenceLength),
			},
		),
	)

	if e != nil {
		return nil, fmt.Errorf("load tokenizer: %w", e)
	}

	counter, f := newCounter(path)

	if f != nil {
		errors.PanicOnError(tokenizer.Close())

		return nil, fmt.Errorf("load counter: %w", f)
	}

	result := &Reranker{
		name:           name,
		sequenceLength: sequenceLength,
		tokenizer:      tokenizer,
		counter:        counter,
	}

	if g := result.measureSpecials(); g != nil {
		errors.PanicOnError(result.Close())

		return nil, fmt.Errorf("measure special tokens: %w", g)
	}

	session, h := newSession(
		join.Join(directory, constant.RerankModelFile),
		sequenceLength,
	)

	if h != nil {
		errors.PanicOnError(result.Close())

		return nil, h
	}

	result.session = session

	return result, nil
}
