package revision

import "go/token"

func closesGroup(t token.Token) bool {
	return t == token.RPAREN || t == token.RBRACE || t == token.RBRACK
}
