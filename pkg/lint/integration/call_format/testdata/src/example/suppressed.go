package example

func suppressed() {
	twoArguments(
		"something-long-enough",
		"to-push-this-line-well-past-the-eighty-character-column-limit-easily",
	) // goanalyze:ignore
}
