package constant

const MemoryWorkflowURI = "gomemoryd://guide/memory-workflow"
const MemoryWorkflow = `# Memory Workflow

gomemoryd stores persistent memories across sessions. Memories
survive session boundaries - what one session saves, the next
session reads.

## Profile

Call profile on your first turn. It returns three tiers:

1. **Always tier** - full content of memories tagged "always".
   These load every session regardless of topic. Use for
   behavioral guidance, conventions, and relationship context.
2. **Topic tier** - full content of memories matching the
   session topic (hybrid search). Relevant context surfaced
   automatically.
3. **Index tier** - brief listing of all other active memories
   (id, description, tags). Scan to decide what to load in
   full via get_memory.

Profile accepts an optional topic parameter for better
topic-tier matching, and an optional scope parameter to
profile a named scope instead of the default one.

## Scopes

The default scope holds session memories - everything above
applies to it. Named scopes hold separate corpora: memories
compiled from document trees (lore, workflows) or any other
isolated cluster. Profile, search, and list accept a scope
parameter; a scoped profile returns that scope's always tier
and index only. The scope names "default" and "all" are
reserved ("all" crosses scopes in search and list).

## Groups

Memories form parent/child clusters. A parent carries a
description (the hallway line) and usually no content; children
carry the content, ordered by ordinal. The profile index
collapses children to name listings under their parent.
get_memory_group fetches a parent with all children in one
call - the standard move for document-sourced clusters. The
response ends with a deduplicated relations section: every
edge touching a group member, outward edges (doors to other
memories) first, internal topology after. set_parent moves a
memory into or out of a cluster.

## Document-sourced memories

Memories carrying a provenance file are compiled from source
documents and are read-only through these tools - update,
forget, and tag answer with the file path to edit instead.
The source document is the edit path; an importer propagates
changes and owns creation and deletion for its scope.

## Creating and updating

save_memory creates a new memory. update_memory updates an
existing one by memory_id (required). Updates record the
previous version - content history is preserved. Tags are
preserved across updates.

Each memory has:
- **content** - the memory text
- **description** - one-line summary for index listings and
  search results
- **type** - categorization: user, feedback, project, reference

## Growth

The corpus grows through the graph. New material lands as
leaves and edges - a new memory parented where it belongs and
related to what it touches - not as paragraphs appended to
existing memories. Appending is the exception, not the default:
it fattens a slot until traversal must consume everything to
find anything. A memory whose name wants commas is two
memories.

The split axis is retrieval, not taxonomy: ask what one moment
needs together, not what one topic contains. Rule-shaped
separation is only one valid shape.

Edits made mid-work accrue lint by nature - a session saving in
the middle of something else cannot see the whole shape. That
is expected, not failure; it is why tending recurs. The
patterns worth acting on only become visible when tending
re-reads the shape whole.

## Tags

Tags organize memories for retrieval. Use tag_memory to add,
remove, or replace tags on a memory.

Two tags have special meaning:

- **always** - always-tagged memories load in full on every
  profile call. Use sparingly for content that every session
  needs.
- **no-index** - removes a memory from the profile index tier.
  The memory stays reachable through topic matching, search,
  and relations. Use for depth leaves that should be found by
  relevance instead of occupying index space. A no-index
  memory with no relations and no parent is a hidden memory -
  it only surfaces when something searches for it specifically.

All other tags are freeform strings. No fixed vocabulary -
create tags that serve the retrieval patterns you need.

## Search and retrieval

search_memories is keyword full-text search - use single
keywords, not sentences; a sentence must match every word and
returns an empty result. For semantic sentence-shaped queries,
use goqueryd (see Search index below). Optional type, tag, and
scope filters. Hits carry their scope and their parent's name -
a shard hit names its document.

get_memory retrieves a memory by ID with full content and its
relations. Pass include_history to see previous versions, or
memory_ids (comma-separated) to fetch several memories in one
call - the move for following a related list without a call
per neighbor.

list_memories shows all active memories with optional type,
tag, and scope filters.

Read tools return a compact projection - content, tags,
metadata, ordinals - without timestamps or provenance. Pass
detail=true on get_memory, get_memory_group, or
search_memories for the full record.

## Relations

relate_memories creates a bidirectional link between two
memories. get_memory returns each neighbor under "related" -
identifier, name, description, type, and tags. Use for
memories that inform each other but shouldn't be merged.

Relations carry an optional type from a fixed vocabulary:

- **affinity** - thematically adjacent, co-occurring topics
  (undirected)
- **informs** - rules or patterns that temper or apply with
  each other (undirected)
- **grounds** - an abstraction pointing at its living referent
  (source grounds in target)
- **mechanism** - causal or physical dependency between
  components (source acts on target)
- **sequence** - ordered succession (source precedes target)
- **deep-dive** - a summary pointing where to expand for depth
  (source expands into target)

Untyped relations are legal and remain so. Relating an
existing pair with a type retypes the edge; relating without
one leaves an existing type alone. unrelate_memories removes
an edge - the stored row is directional, and removal is
permanent with no history.

Relations enable incremental deepening: land on a memory via
the index or search, read its edges, follow the ones that
pull. Search teleports you to a door; relations tell you
which rooms adjoin. A neighbor's tags may point at a whole
tag category worth listing, not just the one memory.

## Deletion

forget_memory soft-deletes a memory. The memory is marked
inactive and a version record is created. Content is preserved
in history but the memory no longer appears in profile, search,
or list results.

## Search index

Memories are pushed to goqueryd on save and delete. This means
memories are searchable alongside documents, session summaries,
and code via goqueryd's hybrid search - semantic, so full
sentences work there. Each scope projects into a goqueryd
collection of the same name; the default scope projects into
the "memories" collection. Source type is "memory".
`
