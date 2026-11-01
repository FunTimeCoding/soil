package constant

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/types/embed_model"

var (
	NomicEmbedText = embed_model.New(
		"nomic-embed-text",
		"search_document: ",
		"search_query: ",
	)
	NomicEmbedTextStudio = embed_model.New(
		"text-embedding-nomic-embed-text-v1.5@f16",
		"search_document: ",
		"search_query: ",
	)
	Qwen3Embedding = embed_model.New(
		"qwen3-embedding:0.6b",
		"",
		"Instruct: Given a web search query, retrieve relevant passages that answer the query\nQuery: ",
	)
)

var EmbedModels = []*embed_model.Model{
	NomicEmbedText,
	NomicEmbedTextStudio,
	Qwen3Embedding,
}
