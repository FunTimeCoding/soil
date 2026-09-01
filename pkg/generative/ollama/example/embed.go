package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Embed() {
	o := ollama.NewEnvironment()

	for _, e := range o.MustEmbedding(
		constant.Llama31,
		"What are embeddings?",
	) {
		console.Format("Embedding: %+v\n", e)
	}
}
