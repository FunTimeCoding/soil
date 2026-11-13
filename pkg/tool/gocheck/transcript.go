package gocheck

import (
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreach"
)

func Transcript(reach string) (string, error) {
	var session []string

	if reach != constant.OwnSession {
		session = []string{reach}
	}

	return goreach.TranscriptPath("", session)
}
