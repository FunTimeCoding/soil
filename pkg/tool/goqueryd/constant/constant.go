package constant

import "github.com/funtimecoding/soil/pkg/identity"

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
	EmbedModel       = "nomic-embed-text"
	ExpandModel      = ""
	DefaultGlob      = "**/*.md"
	RrfK             = 60

	ExpandPrompt = `/no_think Expand this search query into 2-4 alternative search variants. Output each on its own line in the format "type: query" where type is one of:
- lex: keyword variant for full-text search (specific terms, synonyms, abbreviations)
- vec: semantic variant for vector similarity (rephrase the meaning)
- hyde: hypothetical document snippet that would answer the query

Example input: "how does auth work"
Example output:
lex: authentication authorization login
vec: user authentication flow and session management
hyde: The authentication system uses JWT tokens issued on login. Sessions are stored in Redis with a 24-hour TTL.

Query: `

	DashboardTitle   = "Dashboard"
	DashboardPath    = "/"
	SearchTitle      = "Search"
	SearchPath       = "/search"
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
