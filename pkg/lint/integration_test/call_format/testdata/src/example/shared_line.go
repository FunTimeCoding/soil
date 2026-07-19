package example

func sharedLineArgs() {
	withStruct(
		"name", Options{
			Value: 1,
		},
	)
}
