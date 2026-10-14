package system

import (
	"archive/tar"
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func TarSegment(write func(*tar.Writer)) []byte {
	var b bytes.Buffer
	t := tar.NewWriter(&b)
	write(t)
	errors.PanicClose(t)

	return b.Bytes()[:b.Len()-constant.TarFooterSize]
}
