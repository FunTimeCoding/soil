package bravo

import "example/pkg/alfa"

type Bravo struct {
	*alfa.Alfa
	Named *alfa.Alfa
}
