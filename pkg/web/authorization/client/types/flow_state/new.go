package flow_state

func New(
	verifier string,
	state string,
	returnPath string,
	callbackLocator string,
) *State {
	return &State{
		Verifier:        verifier,
		State:           state,
		ReturnPath:      returnPath,
		CallbackLocator: callbackLocator,
	}
}
