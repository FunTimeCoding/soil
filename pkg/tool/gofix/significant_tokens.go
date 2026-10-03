package gofix

import (
	"go/scanner"
	"go/token"
)

func significantTokens(source []byte) ([]int, []int) {
	fileSet := token.NewFileSet()
	file := fileSet.AddFile("", fileSet.Base(), len(source))
	var s scanner.Scanner
	s.Init(file, source, nil, 0)
	var offsets []int
	var lines []int

	for {
		position, kind, _ := s.Scan()

		if kind == token.EOF {
			return offsets, lines
		}

		if kind == token.COMMA || kind == token.SEMICOLON {
			continue
		}

		offsets = append(offsets, file.Offset(position))
		lines = append(lines, file.Line(position))
	}
}
