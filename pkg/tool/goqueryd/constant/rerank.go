package constant

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_model"

const (
	RerankDirectoryEnvironment = "RERANK_DIRECTORY"
	RerankModelEnvironment     = "RERANK_MODEL"
	RerankSequenceEnvironment  = "RERANK_SEQUENCE"
	RerankModelFile            = "model.onnx"
	RerankTokenizerFile        = "tokenizer.json"
	RerankQueryReserve         = 32
	EmbedTokenizerEnvironment  = "EMBED_TOKENIZER"
	EmbedTokenizerDefault      = "~/.local/share/soil/nomic-embed-text/tokenizer.json"
	EmbedTokenizerSource       = "https://huggingface.co/nomic-ai/nomic-embed-text-v1.5/resolve/main/tokenizer.json"
	EmbedTokenizerUnset        = "%s not set - fetch nomic's tokenizer: curl -L --create-dirs -o %s %s"
	EmbedTokenizerUnreadable   = "%s: %w - fetch nomic's tokenizer: curl -L --create-dirs -o %s %s"
	EmbedWindow                = 2048
	EmbedOverWindow            = "over the embedder's window: %d tokens against %d"
	NoChunkPosition            = -1
	MockRerankerName           = "mock"
)

var (
	BgeRerankerBase           = rerank_model.New("bge-reranker-base", 512, 512)
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
