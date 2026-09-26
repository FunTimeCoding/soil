# Chunking Strategy

Long documents are split at natural boundaries. Markdown
files break at headings and paragraph gaps. Go source files
break at top-level declarations.

Each chunk overlaps with its neighbors to preserve context
across boundaries. Code fences are never split mid-block.
