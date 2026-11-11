package lib

type Token struct {
	Value string
}

func Next() *Token {
	return &Token{}
}
