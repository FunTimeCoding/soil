package example

func DroppedChain() {
	newClient().Work()
}

func DroppedBuiltChain() {
	newClient().With().Work()
}
