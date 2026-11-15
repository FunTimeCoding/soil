package repository_tester

import "slices"

func Identified(
	directory string,
	arguments ...string,
) string {
	return Command(
		directory,
		slices.Concat(
			[]string{
				"-c",
				"user.name=alfa",
				"-c",
				"user.email=alfa@example.org",
			},
			arguments,
		)...,
	)
}
