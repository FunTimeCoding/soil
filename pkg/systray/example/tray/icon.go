//go:build local

package tray

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"image"
	"image/color"
	"image/png"
)

func icon() []byte {
	size := 22
	m := image.NewRGBA(image.Rect(0, 0, size, size))

	for y := range size {
		for x := range size {
			m.SetRGBA(x, y, color.RGBA{R: 46, G: 160, B: 67, A: 255})
		}
	}

	var b bytes.Buffer
	errors.PanicOnError(png.Encode(&b, m))

	return b.Bytes()
}
