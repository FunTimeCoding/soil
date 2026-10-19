package error_answer

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path"
	"strings"
)

func knownClass(location string) string {
	if !strings.Contains(location, constant.ErrorClassSegment) {
		return ""
	}

	class := path.Base(location)

	if _, known := errorsConstant.Answers[class]; !known {
		return ""
	}

	return class
}
