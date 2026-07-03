package tab

type Tab struct {
	Description    string `json:"description"`
	Frontend       string `json:"devtoolsFrontendUrl"`
	Identifier     string `json:"id"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Locator        string `json:"url"`
	WebSocketDebug string `json:"webSocketDebuggerUrl"`
	FaviconLocator string `json:"faviconUrl,omitempty"`
	// ParentIdentifier is set on iframe targets and holds the
	// identifier of the tab the frame renders in.
	ParentIdentifier string `json:"parentId,omitempty"`
}
