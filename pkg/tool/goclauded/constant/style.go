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
	.status-disconnected { background: #ff4136; }
	.kind-badge {
		font-size: 0.75rem;
		padding: 0.1rem 0.4rem;
		border-radius: 4px;
		font-weight: bold;
	}
	.kind-complete { background: #2ecc40; color: #fff; }
	.kind-update { background: #0074d9; color: #fff; }
	table { margin-bottom: 0; }
`
