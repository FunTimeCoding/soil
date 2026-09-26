# Document Indexing

Collections scan directories for files matching a glob pattern.
Each file is hashed, chunked, and inserted into the full text
search index. Changed files are re-indexed on the next scan.

Push collections receive documents directly via the API
instead of scanning the filesystem.
