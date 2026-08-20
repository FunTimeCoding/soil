package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
)

var Identity = identity.New(
	"goqueryd",
	"Local search engine with hybrid BM25 + vector + cross-encoder reranking",
	"goqueryd",
).WithInstructions(
	"Local search engine with hybrid BM25 + vector + cross-encoder reranking. Read the goqueryd://guide/search-workflow resource for collection types, search pipeline, source types, and index management.",
)

const (
	Search           = "search"
	Status           = "status"
	Get              = "get"
	List             = "list"
	Push             = "push"
	Delete           = "delete"
	Tag              = "tag"
	Index            = "index"
	Embed            = "embed"
	AddCollection    = "add_collection"
	DeleteCollection = "delete_collection"
	ListTags         = "list_tags"
	AddContext       = "add_context"
	RemoveContext    = "remove_context"
	ListContexts     = "list_contexts"
	ListMetadata     = "list_metadata"

	Metadata    = "metadata"
	Collection  = "collection"
	Mode        = "mode"
	Path        = "path"
	Full        = "full"
	SourceType  = "source_type"
	Body        = "body"
	Pattern     = "pattern"
	PathPrefix  = "path_prefix"
	Description = "description"
	Key         = "key"

	DefaultSequenceLength = 512
	ModelEnvironment      = "RERANK_MODEL"
	TokenizerEnvironment  = "RERANK_TOKENIZER"

	ChunkSize        = 3600
	ChunkOverlap     = 540
	ChunkWindow      = 800
	SnippetMaxLength = 400
	ListPage         = 500
	DefaultGlob      = "**/*.md"
	RrfK             = 60

	DashboardTitle   = "Dashboard"
	DashboardPath    = "/"
	SearchTitle      = "Search"
	CollectionsTitle = "Collections"
	CollectionsPath  = "/collections"
	Identifier       = "identifier"

	FixtureAuthorKey  = "author"
	FixtureTagKey     = "tag"
	FixtureBuildValue = "build"

	TestBody = `# Search Pipeline

This document describes the hybrid search pipeline.

## Keyword Matching

BM25 scores documents by term frequency.
Rare terms receive higher weight than common ones.

## Vector Similarity

Embeddings capture semantic meaning beyond exact terms.
Cosine distance measures how close two vectors are.

## Cross-Encoder Reranking

A cross-encoder scores each query-document pair directly.
It reranks the merged candidate set by relevance.`
)

var FencePattern = regexp.MustCompile(join.Empty(`\n`, "```"))

var (
	HyphenatedPattern = regexp.MustCompile(
		`^[\p{L}\p{N}][\p{L}\p{N}'-]*-[\p{L}\p{N}][\p{L}\p{N}'-]*$`,
	)
	HeadingPattern  = regexp.MustCompile(`(?m)^##?\s+(.+)$`)
	NonAlphanumeric = regexp.MustCompile(`[^\p{L}\p{N}'_]+`)
)
