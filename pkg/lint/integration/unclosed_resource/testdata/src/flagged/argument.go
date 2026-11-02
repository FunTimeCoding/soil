package example

func use(*Client) {}

func PassedFresh() {
	use(newClient())
}
