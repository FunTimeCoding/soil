package view

import "strings"

func New(text string) *View {
	return &View{Text: text, Lines: strings.Count(text, "\n")}
}
