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
	html, body {
		height: 100%;
		margin: 0;
		overflow: hidden;
	}
	.conversation-layout {
		display: grid;
		grid-template-columns: 350px 1fr;
		height: 100vh;
	}
	.sidebar {
		overflow-y: auto;
		border-right: 1px solid var(--pico-muted-border-color);
		padding: 0.5rem;
	}
	.sidebar-entry {
		padding: 0.5rem;
		cursor: pointer;
		border-radius: var(--pico-border-radius);
		margin-bottom: 0.25rem;
	}
	.sidebar-entry:hover {
		background: var(--pico-muted-border-color);
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
		opacity: 0;
		cursor: pointer;
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
	.search-input {
		width: 100%;
		padding: 0.4rem 0.5rem;
		margin-bottom: 0.25rem;
		font-size: 0.85rem;
	}
	.search-kinds {
		display: flex;
		gap: 0.75rem;
		font-size: 0.8rem;
	}
	.search-kinds label {
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
		padding: 0.5rem;
		cursor: pointer;
		border-radius: var(--pico-border-radius);
	}
	.search-result summary:hover {
		background: var(--pico-muted-border-color);
	}
	.search-result summary small {
		color: var(--pico-muted-color);
		display: block;
	}
	.search-snippet {
		padding: 0.35rem 0.5rem 0.35rem 1.25rem;
		font-size: 0.85rem;
		cursor: pointer;
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
