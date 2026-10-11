package base

import "sync"

var sharedReranker = sync.OnceValue(newReranker)
