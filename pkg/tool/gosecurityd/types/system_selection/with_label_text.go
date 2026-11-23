package system_selection

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strings"
)

func (s *Selection) WithLabelText(text string) (*Selection, error) {
	key, value, _ := strings.Cut(text, constant.LabelSeparator)

	if key == "" || value == "" {
		return nil, validation.New("label filter malformed: %s", text)
	}

	return s.WithLabel(key, value), nil
}
