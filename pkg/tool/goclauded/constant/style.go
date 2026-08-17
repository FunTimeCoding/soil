package constant

const InlineStyle = `
	pre { white-space: pre-wrap; word-break: break-word; }
	.roster-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
		gap: 1rem;
		margin-bottom: 1rem;
	}
	a.session-card {
		display: block;
		text-decoration: none;
		color: inherit;
		border: 1px solid var(--pico-muted-border-color);
		border-radius: var(--pico-border-radius);
		padding: 1rem 1.25rem;
	}
	a.session-card:hover {
		border-color: var(--pico-primary);
	}
	.session-card h4 {
		margin-bottom: 0.25rem;
	}
	.session-card p {
		margin-bottom: 0.25rem;
	}
	.session-card small {
		color: var(--pico-muted-color);
	}
	.status-dot {
		display: inline-block;
		width: 8px;
		height: 8px;
		border-radius: 50%;
		margin-right: 0.4rem;
		vertical-align: middle;
	}
	.status-active { background: #2ecc40; }
	.status-stale { background: #ffdc00; }
	.kind-badge {
		font-size: 0.75rem;
		padding: 0.1rem 0.4rem;
		border-radius: 4px;
		font-weight: bold;
	}
	.kind-complete { background: #2ecc40; color: #fff; }
	.kind-update { background: #0074d9; color: #fff; }
	.session-card table { margin-bottom: 0; }
	.label-pip {
		display: inline-block;
		font-size: 0.75rem;
		padding: 0.15rem 0.5rem;
		margin-right: 0.35rem;
		border-radius: var(--pico-border-radius);
		border: 1px solid var(--pico-muted-border-color);
		background: var(--pico-card-sectioning-background-color);
	}
	.label-key {
		color: var(--pico-muted-color);
		margin-right: 0.3rem;
	}
`
const ConversationStyle = `
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
	.rename-input {
		width: 100%;
		font-size: 0.9rem;
		padding: 0.2rem 0.4rem;
		margin: 0;
	}
	.panel {
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
		border-left: 3px solid #0074d9;
	}
	.message-assistant {
		border-left: 3px solid #2ecc40;
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
	.sidebar-filter {
		width: 100%;
		padding: 0.4rem 0.5rem;
		margin-bottom: 0.5rem;
		font-size: 0.85rem;
	}
`
