package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
)

func Embed() {
	o := ollama.NewEnvironment()

	for _, e := range o.MustEmbedding(
		constant.Llama31,
		"What are embeddings?",
	) {
		fmt.Printf("Embedding: %+v\n", e)
	}
}
