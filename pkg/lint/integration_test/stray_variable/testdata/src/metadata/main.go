package main

var (
	Version   string
	GitHash   string
	BuildDate string
)

var stray = "flagged"

func main() {
	_ = Version + GitHash + BuildDate + stray
}
