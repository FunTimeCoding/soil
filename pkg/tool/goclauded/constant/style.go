package constant

import "github.com/funtimecoding/soil/pkg/strings/join"

const InlineStyle = `
	pre { white-space: pre-wrap; word-break: break-word; }
	.detail-head { display: grid; grid-template-columns: minmax(0, 36rem) minmax(0, 1fr); gap: 2rem; align-items: start; }
`

var ConversationStyle = join.Empty(
	ConversationLayoutStyle,
	ConversationMessageStyle,
	ConversationSearchStyle,
	ConversationHitStyle,
)

const ConversationLayoutStyle = `
	.conversation-layout {
		display: grid;
		grid-template-columns: 350px 1fr;
	}
	.sidebar {
		overflow-y: auto;
		border-right: 1px solid var(--pico-muted-border-color);
		padding: 0.5rem;
	}
	.sidebar-entry {
		position: relative;
		padding: 0.5rem;
		border-radius: var(--pico-border-radius);
		margin-bottom: 0.25rem;
	}
	.sidebar-entry:hover,
	.sidebar-entry[aria-current] {
		background: var(--pico-muted-border-color);
	}
	.entry-link {
		color: inherit;
	}
	.entry-link::after {
		content: "";
		position: absolute;
		inset: 0;
	}
	.sidebar-more {
		color: var(--pico-muted-color);
	}
	.sidebar-entry small {
		color: var(--pico-muted-color);
		display: block;
	}
	.entry-name {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.rename-icon {
		position: relative;
		z-index: 1;
		opacity: 0;
		font-size: 0.85rem;
	}
	.sidebar-entry:hover .rename-icon {
		opacity: 0.5;
	}
	.rename-icon:hover {
		opacity: 1 !important;
	}
`

const ConversationMessageStyle = `	.panel {
		overflow-y: auto;
		padding: 1rem 2rem;
	}
	.panel-placeholder {
		color: var(--pico-muted-color);
		text-align: center;
		margin-top: 40vh;
	}
	.message {
		margin-bottom: 1rem;
		padding: 0.75rem 1rem;
		border-radius: var(--pico-border-radius);
		max-width: 1000px;
		margin-left: auto;
		margin-right: auto;
	}
	.message-user {
		border-left: 3px solid var(--busy);
	}
	.message-assistant {
		border-left: 3px solid var(--okay);
		background: var(--pico-card-background-color);
	}
	.message-role {
		font-size: 0.75rem;
		font-weight: bold;
		text-transform: uppercase;
		color: var(--pico-muted-color);
		margin-bottom: 0.25rem;
	}
	.message-text {
		white-space: pre-wrap;
		word-wrap: break-word;
	}
`

const ConversationSearchStyle = `	.search-form {
		margin-bottom: 0.5rem;
	}
	.search-form fieldset {
		display: flex;
		flex-wrap: wrap;
		gap: 0.75rem;
		font-size: 0.8rem;
	}
	.search-form fieldset legend {
		width: 100%;
	}
	.search-form fieldset label {
		margin: 0;
	}
	.search-progress {
		display: block;
		color: var(--pico-muted-color);
		margin-bottom: 0.5rem;
	}
	.search-result {
		margin-bottom: 0.25rem;
	}
	.search-result summary {
		position: relative;
		padding: 0.5rem 2rem 0.5rem 0.5rem;
		cursor: pointer;
		border-radius: var(--pico-border-radius);
	}
	.search-result summary::after {
		position: absolute;
		right: 0.5rem;
		top: 0.75rem;
		margin: 0;
	}
	.search-result summary:hover {
		background: var(--pico-muted-border-color);
	}
	.search-result summary small {
		color: var(--pico-muted-color);
		display: block;
	}
	.search-snippet {
		display: block;
		color: inherit;
		padding: 0.35rem 0.5rem 0.35rem 1.25rem;
		font-size: 0.85rem;
		border-radius: var(--pico-border-radius);
		word-wrap: break-word;
	}
	.search-snippet:hover {
		background: var(--pico-muted-border-color);
	}
	.search-snippet small {
		color: var(--pico-muted-color);
		margin-right: 0.4rem;
		text-transform: uppercase;
	}
`

const ConversationHitStyle = `	.message-tool {
		border-left: 3px solid var(--pico-muted-color);
		font-size: 0.9rem;
	}
	.search-hit {
		outline: 1px solid var(--pico-mark-background-color);
	}
	.search-current {
		outline: 3px solid var(--pico-mark-background-color);
	}
	.hit-counter {
		position: sticky;
		bottom: 0;
		text-align: right;
		color: var(--pico-muted-color);
		font-size: 0.85rem;
		padding: 0.25rem 0;
		background: var(--pico-background-color);
	}
`
