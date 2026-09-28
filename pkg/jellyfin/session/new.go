package session

func New(
	identifier string,
	deviceName string,
	deviceIdentifier string,
	client string,
	userName string,
	nowPlayingName string,
	nowPlayingType string,
	nowPlayingIdentifier string,
	playState string,
	position string,
	supportsRemote bool,
) *Session {
	return &Session{
		Identifier:           identifier,
		DeviceName:           deviceName,
		DeviceIdentifier:     deviceIdentifier,
		Client:               client,
		UserName:             userName,
		NowPlayingName:       nowPlayingName,
		NowPlayingType:       nowPlayingType,
		NowPlayingIdentifier: nowPlayingIdentifier,
		PlayState:            playState,
		Position:             position,
		SupportsRemote:       supportsRemote,
	}
}
