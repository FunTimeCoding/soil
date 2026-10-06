package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"testing"
)

func TestNewModelRefusesWindowBeyondMaximum(t *testing.T) {
	t.Setenv(constant.RerankSequenceEnvironment, "1024")
	_, e := rerank.NewModel(constant.BgeRerankerBase)
	assert.Error(t, e)
	assert.String(
		t,
		"window 1024 exceeds bge-reranker-base maximum of 512",
		e.Error(),
	)
}
