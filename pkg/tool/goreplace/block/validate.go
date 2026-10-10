package block

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func (b *Block) validate() error {
	if b.Search == "" {
		return fmt.Errorf(constant.EmptySearch, b.Number)
	}

	switch b.Opening {
	case constant.SpanMarker:
		if b.Closing == constant.ToEndMarker && b.End != "" {
			return fmt.Errorf(constant.TextAfterToEnd, b.Number)
		}

		if b.Closing != constant.ToEndMarker && b.End == "" {
			return fmt.Errorf(constant.EmptyEnd, b.Number, b.Closing)
		}
	case constant.AfterMarker, constant.BeforeMarker:
		if strings.Contains(b.Search, "\n") {
			return fmt.Errorf(constant.MultilineAnchor, b.Number)
		}

		if b.Replace == "" {
			return fmt.Errorf(constant.EmptyInsert, b.Number)
		}
	}

	return nil
}
