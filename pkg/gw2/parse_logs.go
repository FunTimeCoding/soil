package gw2

import (
	"bytes"
	"github.com/dimchansky/utfbom"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func ParseLogs(
	s []byte,
	verbose bool,
) []*response.Log {
	reader, encoding := utfbom.Skip(bytes.NewReader(s))

	if verbose {
		console.Format("Detected encoding: %s\n", encoding)
	}

	var f response.LogFile
	notation.MustDecode(string(system.ReadAll(reader)), &f, true)
	var result []*response.Log

	for k, v := range f.LogsByFilename {
		if v == nil {
			errors.Warning("no data: %s", k)

			continue
		}

		var l response.Log
		errors.PanicOnError(notation.Decode(notation.Encode(v, false), &l))
		result = append(result, &l)
	}

	return result
}
