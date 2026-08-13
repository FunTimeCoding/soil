package constant

const SearchWorkflowURI = "goqueryd://guide/search-workflow"
const SearchWorkflow = `# Search Workflow

goqueryd is a local search engine for documents and code. It
combines BM25 full-text search, vector similarity, and
cross-encoder reranking to return ranked results.

## Collections

Two types of collections hold documents:

**Filesystem collections** scan a directory with a glob pattern
(default ` + "`**/*.md`" + `). Populated by calling index, which
detects new, changed, and removed files via content hashing.

**Push collections** receive documents via the push tool. No
filesystem path - content is sent directly. Used by other
services (gomemoryd pushes memories, goclauded pushes session
summaries). The collection is created implicitly on first push.

Call status to see all collections with document counts and
embedding statistics.

## Searching

The search tool runs a hybrid pipeline:

1. BM25 keyword matching (FTS5)
2. Vector similarity (cosine distance over embeddings)
3. Reciprocal Rank Fusion to merge ranked lists
4. Cross-encoder reranking (bge-reranker-base) for final scoring

Scores above 0.5 are relevant. Scores below 0.1 are noise.

Parameters:
- query (required) - search terms or natural language
- collection - restrict to one collection (omit for all)
- source_type - filter by source type tag (e.g. session-summary,
  memory, seed, design-doc)
- mode - hybrid (default) or keyword (skips vector + reranking)
- full - include full document body in results (default false)
- limit - max results (default 10)

When the embedding backend is unavailable, search degrades to
keyword-only automatically.

## Retrieving documents

Use get with a document path to retrieve full content. Accepts
virtual paths (qmd://collection/path) or relative paths
(collection/path). On miss, suggests similar paths via fuzzy
matching.

Use list to see all documents in a collection with their virtual
paths and titles.

## Source types

Documents can be tagged with a source type for filtering.
Source types are set per path prefix via the tag tool, or per
document on push. Examples: session-summary, memory, seed,
design-doc.

Source type resolution uses most-specific-wins: a document at
plan/seed/goqueryd.md tagged under plan/seed/ as seed returns
seed, not a parent plan/ tag.

## Contexts

Path prefix contexts attach descriptions to search results.
Use add_context to describe what a directory contains - the
description appears on every result from that path. Contexts
are hierarchical: a root context applies to all documents,
more specific prefixes add detail.

## Managing the index

- index - re-scan filesystem collections for changes
- embed - generate embeddings for documents missing them
- add_collection - register a new filesystem collection
- push - add or update a document in a push collection
- delete - remove a document and its embeddings

Pushing unchanged content is cheap - content-hash dedup skips
re-embedding when the body hasn't changed.
`
