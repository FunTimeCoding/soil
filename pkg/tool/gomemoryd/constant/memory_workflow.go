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
topic-tier matching.

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

search_memories finds memories by keyword with optional type
and tag filters. Returns ranked results.

get_memory retrieves a single memory by ID with full content.
Pass include_history to see previous versions.

list_memories shows all active memories with optional type and
tag filters.

## Relations

relate_memories creates a bidirectional link between two
memories. get_memory returns each neighbor under "related" -
identifier, name, description, and tags. Use for memories
that inform each other but shouldn't be merged.

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
and code via goqueryd's hybrid search. The memories collection
uses source type "memory".
`
