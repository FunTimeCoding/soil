package constant

import "time"

const (
	ChannelCapability   = "claude/channel"
	ChannelNotification = "notifications/claude/channel"
	ChannelContent      = "content"
	ChannelMeta         = "meta"
	ChannelKindMeta     = "kind"
	ChannelNonceMeta    = "nonce"
	ChannelCallsignMeta = "callsign"
	ChannelAttachKind   = "attach"
	ChannelStalledKind  = "stalled"

	ChannelConfirmTool = "confirm_channel"

	ChannelAttachMessage = "channel attached and waiting. Confirm with this event's nonce and your coordination callsign to open delivery."

	ChannelConfirmDescription = "Confirm that a channel attach event reached this session, opening delivery. Call only after receiving the event."

	ChannelNonceDescription = "The nonce carried by the attach event."

	ChannelCallsignDescription = "This session's coordination callsign, which names the queue to deliver."

	ChannelNonceMismatch = "nonce mismatch - delivery stays closed. Pass the nonce from the attach event this session received."

	ChannelCallsignRequired = "callsign is required - it names the queue this channel delivers."

	ChannelCallsignMismatch = "callsign mismatch - delivery stays closed. This channel resolved %q for its session and you passed %q, so one of the two is bound to the wrong session."

	ChannelOpened = "channel open - coordination traffic now arrives here instead of the pre-prompt context"

	ChannelStallThreshold = 3

	ChannelAttachInterval = time.Minute

	ChannelAttachMaximum = 5 * time.Minute

	ChannelStalledMessage = "channel delivery is stalled: %d consecutive poll failures reaching the coordination daemon. Nothing is being delivered here, and pending traffic is waiting for your next prompt instead."

	ChannelResumedMessage = "channel delivery resumed."
)
