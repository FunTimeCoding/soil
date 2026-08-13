package relational

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

func (d *Database) validateDatabase(n string) {
	if n == "" {
		log.Panicf("database name cannot be empty")
	}

	if len(n) > 63 {
		log.Panicf("database name cannot be longer than 63 characters")
	}

	if !unicode.IsLetter(rune(n[0])) && n[0] != '_' {
		log.Panicf("database name must start with a letter or underscore")
	}

	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_$]*$`).MatchString(n) {
		log.Panicf(
			"database name can only contain letters, digits, underscores, and dollar signs",
		)
	}

	lower := strings.ToLower(n)

	for _, reserved := range []string{
		"postgres",
		"template0",
		"template1",
		"information_schema",
		"pg_catalog",
	} {
		if lower == reserved {
			log.Panicf("database name cannot be a reserved name: %s", reserved)
		}
	}
}
