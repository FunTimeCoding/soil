package model

type Reader interface {
	Press(key rune)
	Release(key rune)
}
