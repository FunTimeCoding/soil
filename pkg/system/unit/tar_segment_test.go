package unit

import (
	"archive/tar"
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/unit/tar_tester"
	"io"
	"testing"
)

func TestTarSegmentIsArchiveWithoutItsEnd(t *testing.T) {
	var full bytes.Buffer
	w := tar.NewWriter(&full)
	tar_tester.Entry(stringsConstant.LowerAlfa, stringsConstant.LowerAlfa)(w)
	errors.PanicClose(w)
	segment := system.TarSegment(
		tar_tester.Entry(stringsConstant.LowerAlfa, stringsConstant.LowerAlfa),
	)
	assert.Integer(t, full.Len()-1024, len(segment))
	assert.True(t, bytes.Equal(full.Bytes()[:len(segment)], segment))
	assert.True(t, bytes.Equal(make([]byte, 1024), full.Bytes()[len(segment):]))
}

func TestTarSegmentsConcatenateIntoOneArchive(t *testing.T) {
	var b bytes.Buffer
	b.Write(
		system.TarSegment(
			tar_tester.Entry(
				stringsConstant.LowerAlfa,
				stringsConstant.LowerAlfa,
			),
		),
	)
	b.Write(
		system.TarSegment(
			tar_tester.Entry(
				stringsConstant.LowerBravo,
				stringsConstant.LowerBravo,
			),
		),
	)
	r := tar.NewReader(&b)
	var names []string

	for {
		h, e := r.Next()

		if e == io.EOF {
			break
		}

		assert.FatalOnError(t, e)
		names = append(names, h.Name)
	}

	assert.Strings(t, []string{"alfa", "bravo"}, names)
}
