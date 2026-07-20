package main

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/usage"
	anthropic "github.com/funtimecoding/soil/pkg/generative/anthropic/example"
	gguf "github.com/funtimecoding/soil/pkg/generative/gguf/example"
	langchain "github.com/funtimecoding/soil/pkg/generative/langchain/example"
	"github.com/funtimecoding/soil/pkg/generative/langchain/example/function"
	mistral "github.com/funtimecoding/soil/pkg/generative/mistral/example"
	anthropicServer "github.com/funtimecoding/soil/pkg/generative/model_context/example/anthropic"
	"github.com/funtimecoding/soil/pkg/generative/model_context/example/mark"
	ollama "github.com/funtimecoding/soil/pkg/generative/ollama/example"
	openWebUI "github.com/funtimecoding/soil/pkg/generative/open_webui/example"
	openai "github.com/funtimecoding/soil/pkg/generative/openai/example"
)

func main() {
	usage.Usage()

	if false {
		usage.Debug()
		mistral.Prompt()
		ollama.Chat()
		ollama.Benchmark()
		ollama.ClassifyAlert()
		ollama.Fast()
		ollama.Custom()
		ollama.Stream()
		ollama.Simple()
		ollama.Show()
		ollama.List()
		ollama.Embed()
		ollama.Running()
		ollama.Heartbeat()
		openai.Token()
		openai.Official()
		openai.Alternate()
		openai.Client()
		anthropic.Official()
		anthropic.Alternate()
		langchain.Chroma()
		function.Function()
		langchain.Local()
		openWebUI.Load()
		anthropicServer.Run()
		mark.Main()
		gguf.Read()
	}
}
