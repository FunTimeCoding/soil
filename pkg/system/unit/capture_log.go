package unit

import (
	"bytes"
	"log"
)

func captureLog(f func()) string {
	var b bytes.Buffer
	previous := log.Writer()
	log.SetOutput(&b)
	defer log.SetOutput(previous)
	f()

	return b.String()
}
