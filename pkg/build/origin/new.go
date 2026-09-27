package origin

func New(
	hash string,
	time string,
) *Origin {
	return &Origin{Hash: hash, Time: time}
}
