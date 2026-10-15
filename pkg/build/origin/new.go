package origin

func New(
	version string,
	hash string,
	time string,
) *Origin {
	return &Origin{Version: version, Hash: hash, Time: time}
}
