package gohook_tester

import "sync"

var binaryDirectory = sync.OnceValue(buildBinary)
