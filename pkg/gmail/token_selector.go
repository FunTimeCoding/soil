package gmail

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/maps"
	"golang.org/x/oauth2"
	"log"
)

func tokenSelector(m map[string]*oauth2.Token) *oauth2.Token {
	fmt.Printf("Select token by number\n")
	sorted := maps.StringKeys(m)

	for i, k := range sorted {
		fmt.Printf("%d: %s\n", i+1, k)
	}

	var choice int
	_, e := fmt.Scanf("%d", &choice)
	errors.PanicOnError(e)

	if choice < 1 || choice > len(m) {
		log.Panicf("invalid: %d", choice)
	}

	result := m[sorted[choice-1]]

	if result.Valid() || result.RefreshToken != "" {
		return result
	}

	return nil
}
