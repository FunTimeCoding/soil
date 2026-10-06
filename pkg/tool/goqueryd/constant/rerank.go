package constant

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_model"

const (
	RerankDirectoryEnvironment = "RERANK_DIRECTORY"
	RerankModelEnvironment     = "RERANK_MODEL"
	RerankSequenceEnvironment  = "RERANK_SEQUENCE"
	RerankModelFile            = "model.onnx"
	RerankTokenizerFile        = "tokenizer.json"
	RerankQueryReserve         = 32
	NoChunkPosition            = -1
	MockRerankerName           = "mock"
)

var (
	BgeRerankerBase = rerank_model.New("bge-reranker-base", 512, 512)
	GteRerankerModernBertBase = rerank_model.New(
		"gte-reranker-modernbert-base",
		512,
		8192,
	)
)

var RerankModels = []*rerank_model.Model{
	BgeRerankerBase,
	GteRerankerModernBertBase,
}
