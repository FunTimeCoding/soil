package goflow

import (
	"errors"
	"github.com/yuin/goldmark/v2/ast"
	"slices"
)

func verify(
	source []byte,
	document ast.Node,
	rewrapped []byte,
) error {
	after := parse(rewrapped)

	if !slices.Equal(words(source, document), words(rewrapped, after)) {
		return errors.New("word sequence changed")
	}

	if !slices.Equal(shape(document), shape(after)) {
		return errors.New("document structure changed")
	}

	return nil
}
