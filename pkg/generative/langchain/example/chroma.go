package example

// noinspection SpellCheckingInspection
func Chroma() {
	// TODO: Another import hell
	/*
	   import (
	   	"context"
	   	"fmt"
	   	"github.com/funtimecoding/soil/pkg/errors"
	   	"github.com/google/uuid"
	   	"github.com/tmc/langchaingo/embeddings"
	   	"github.com/tmc/langchaingo/llms/ollama"
	   	"github.com/tmc/langchaingo/schema"
	   	"github.com/tmc/langchaingo/vectorstores"
	   	"github.com/tmc/langchaingo/vectorstores/chroma"
	   	"os"
	   	"strings"
	   )
	   func Chroma() {
	   	c, clientFail := ollama.New(ollama.WithModel(constant.Llama31))
	   	errors.PanicOnError(clientFail)
	   	embedder, embedderFail := embeddings.NewEmbedder(c)
	   	errors.PanicOnError(embedderFail)
	   	store, storeFail := chroma.New(
	   		chroma.WithChromaURL(os.Getenv("CHROMA_LOCATOR")),
	   		chroma.WithEmbedder(embedder),
	   		chroma.WithDistanceFunction("cosine"),
	   		chroma.WithNameSpace(uuid.New().String()),
	   	)
	   	errors.PanicOnError(storeFail)
	   	type meta = map[string]any
	   	x := context.Background()
	   	_, addFail := store.AddDocuments(
	   		x,
	   		[]schema.Document{
	   			{
	   				PageContent: "Tokyo",
	   				Metadata:    meta{"population": 9.7, "area": 622},
	   			},
	   			{
	   				PageContent: "Kyoto",
	   				Metadata:    meta{"population": 1.46, "area": 828},
	   			},
	   			{
	   				PageContent: "Hiroshima",
	   				Metadata:    meta{"population": 1.2, "area": 905},
	   			},
	   			{
	   				PageContent: "Kazuno",
	   				Metadata:    meta{"population": 0.04, "area": 707},
	   			},
	   			{
	   				PageContent: "Nagoya",
	   				Metadata:    meta{"population": 2.3, "area": 326},
	   			},
	   			{
	   				PageContent: "Toyota",
	   				Metadata:    meta{"population": 0.42, "area": 918},
	   			},
	   			{
	   				PageContent: "Fukuoka",
	   				Metadata:    meta{"population": 1.59, "area": 341},
	   			},
	   			{
	   				PageContent: "Paris",
	   				Metadata:    meta{"population": 11, "area": 105},
	   			},
	   			{
	   				PageContent: "London",
	   				Metadata:    meta{"population": 9.5, "area": 1572},
	   			},
	   			{
	   				PageContent: "Santiago",
	   				Metadata:    meta{"population": 6.9, "area": 641},
	   			},
	   			{
	   				PageContent: "Buenos Aires",
	   				Metadata:    meta{"population": 15.5, "area": 203},
	   			},
	   			{
	   				PageContent: "Rio de Janeiro",
	   				Metadata:    meta{"population": 13.7, "area": 1200},
	   			},
	   			{
	   				PageContent: "Sao Paulo",
	   				Metadata:    meta{"population": 22.6, "area": 1523},
	   			},
	   		},
	   	)
	   	errors.PanicOnError(addFail)
	   	type exampleCase struct {
	   		name         string
	   		query        string
	   		numDocuments int
	   		options      []vectorstores.Option
	   	}
	   	type filter = map[string]any
	   	cases := []exampleCase{
	   		{
	   			name:         "Up to 5 Cities in Japan",
	   			query:        "Which of these are cities are located in Japan?",
	   			numDocuments: 5,
	   			options: []vectorstores.Option{
	   				vectorstores.WithScoreThreshold(0.8),
	   			},
	   		},
	   		{
	   			name:         "A City in South America",
	   			query:        "Which of these are cities are located in South America?",
	   			numDocuments: 1,
	   			options: []vectorstores.Option{
	   				vectorstores.WithScoreThreshold(0.8),
	   			},
	   		},
	   		{
	   			name:         "Large Cities in South America",
	   			query:        "Which of these are cities are located in South America?",
	   			numDocuments: 100,
	   			options: []vectorstores.Option{
	   				vectorstores.WithFilters(
	   					filter{
	   						"$and": []filter{
	   							{"area": filter{"$gte": 1000}},
	   							{"population": filter{"$gte": 13}},
	   						},
	   					},
	   				),
	   			},
	   		},
	   	}
	   	results := make([][]schema.Document, len(cases))

	   	for ecI, ec := range cases {
	   		docs, searchFail := store.SimilaritySearch(
	   			x,
	   			ec.query,
	   			ec.numDocuments,
	   			ec.options...,
	   		)
	   		errors.PanicOnError(searchFail)
	   		results[ecI] = docs
	   	}

	   	fmt.Printf("Results:\n")

	   	for ecI, ec := range cases {
	   		texts := make([]string, len(results[ecI]))

	   		for docI, doc := range results[ecI] {
	   			texts[docI] = doc.PageContent
	   		}

	   		fmt.Printf("%d. case: %s\n", ecI+1, ec.name)
	   		fmt.Printf("    result: %s\n", join.CommaSpace(texts))
	   	}
	   }
	*/
}
