package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
	"github.com/ollama/ollama/api"
)

type OllamaClient interface {
	Chat(v *chat_request.Request) (*chat_response.Response, error)
	ChatSimple(p string) *chat_response.Response
	MustChat(v *chat_request.Request) *chat_response.Response
	Generate(v *generate_request.Request) (*generate_response.Response, error)
	GenerateSimple(p string) *generate_response.Response
	GenerateFast(p string) *generate_response.Response
	GenerateNotation(p string) *generate_response.Response
	MustGenerate(v *generate_request.Request) *generate_response.Response
	GenerateStream(
		r *api.GenerateRequest,
		output chan<- string,
		response *api.GenerateResponse,
	) error
	MustGenerateStream(
		r *api.GenerateRequest,
		output chan<- string,
		response *api.GenerateResponse,
	)
	Embed(
		x context.Context,
		v []string,
	) ([][]float32, error)
	MustEmbed(
		x context.Context,
		v []string,
	) [][]float32
	EmbedSingle(
		x context.Context,
		v string,
	) ([]float32, error)
	MustEmbedSingle(
		x context.Context,
		v string,
	) []float32
	Embedding(
		model string,
		p string,
	) ([]float64, error)
	MustEmbedding(
		model string,
		p string,
	) []float64
	List() ([]api.ListModelResponse, error)
	MustList() []api.ListModelResponse
	Show(model string) (*api.ShowResponse, error)
	MustShow(model string) *api.ShowResponse
	Running() ([]api.ProcessModelResponse, error)
	MustRunning() []api.ProcessModelResponse
	Version() (string, error)
	MustVersion() string
	Reachable(x context.Context) error
	Locator() string
	Close()
}
