package base

import "sync"

// Shared across the binary's tests - loading costs ~500ms and
// inference is read-only. Never closed; process exit reaps it.
var sharedReranker = sync.OnceValue(newReranker)
