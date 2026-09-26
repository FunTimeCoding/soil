# Search Pipeline

The hybrid search pipeline combines keyword matching with
vector similarity. Documents are ranked by reciprocal rank
fusion after both retrieval methods return candidates.

Query expansion generates alternative phrasings to improve
recall. A cross-encoder reranks the merged candidate set
for final relevance scoring.
