package example

func suppressedByKey() {
	twoArguments(
		"something-long-enough",
		"to-push-this-line-well-past-the-eighty-character-column-limit-easily",
	) // goanalyze:ignore call_format
}
