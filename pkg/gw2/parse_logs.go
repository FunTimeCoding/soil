package gw2

import (
	"bytes"
	"fmt"
	"github.com/dimchansky/utfbom"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func ParseLogs(
	s []byte,
	verbose bool,
) []*log_manager.Log {
	reader, encoding := utfbom.Skip(bytes.NewReader(s))

	if verbose {
		fmt.Printf("Detected encoding: %s\n", encoding)
	}

	var f log_manager.LogFile
	notation.MustDecode(string(system.ReadAll(reader)), &f, true)
	var result []*log_manager.Log

	for k, v := range f.LogsByFilename {
		if v == nil {
			errors.Warning("no data: %s", k)

			continue
		}

		var log log_manager.Log
		errors.PanicOnError(
			notation.Decode(notation.Encode(v, false), &log),
		)
		result = append(result, &log)
	}

	return result
}
