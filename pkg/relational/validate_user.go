package relational

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

func (d *Database) validateUser(n string) {
	if n == "" {
		log.Panicf("user cannot be empty")
	}

	if len(n) > 63 {
		log.Panicf("user cannot be longer than 63 characters")
	}

	if !unicode.IsLetter(rune(n[0])) && n[0] != '_' {
		log.Panicf("user must start with a letter or underscore")
	}

	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_$]*$`).MatchString(n) {
		log.Panicf(
			"user can only contain letters, digits, underscores, and dollar signs",
		)
	}

	lower := strings.ToLower(n)

	for _, reserved := range []string{
		"user",
		"role",
		"group",
		"public",
		"postgres",
		"admin",
		"root",
		"select",
		"insert",
		"update",
		"delete",
		"create",
		"drop",
		"alter",
		"grant",
		"revoke",
		"all",
		"any",
		"some",
		"table",
		"database",
	} {
		if lower == reserved {
			log.Panicf("user cannot be a reserved word: %s", reserved)
		}
	}
}
