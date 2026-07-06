package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gpustack/gguf-parser-go"
)

func Read() {
	a := argument.NewSimple("gguf-read")
	a.ParseSimple()
	result, e := gguf_parser.ParseGGUFFile(
		a.RequiredPositional(0, "PATH"),
		gguf_parser.SkipLargeMetadata(),
		gguf_parser.UseMMap(),
	)
	errors.PanicOnError(e)
	fmt.Printf("Parameters: %+v\n", result.Header)
}
