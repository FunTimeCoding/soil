# Context Resolution

Path prefix contexts attach descriptions to search results.
Contexts are hierarchical - a root context applies to all
documents, and more specific prefixes add detail.

When a document matches a search, all contexts whose prefix
matches the document path are collected and joined. The most
general context appears first.
