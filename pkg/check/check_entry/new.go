package check_entry

import (
	"github.com/funtimecoding/soil/pkg/check/constant"
	"log"
	"slices"
	"time"
)

func New(
	level string,
	text string,
) *Entry {
	if !slices.Contains(constant.Levels, level) {
		log.Panicf("unexpected level: %s", level)
	}

	return &Entry{Time: time.Now(), Level: level, Text: text}
}
