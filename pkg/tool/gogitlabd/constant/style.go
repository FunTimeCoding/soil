package constant

const InlineStyle = `
.status-icon {
	width: 1.15rem;
	height: 1.15rem;
	vertical-align: middle;
}
.board-table td { padding: 0.6rem 0.75rem; }
.board-table .reference { color: var(--pico-secondary); }
.detail-header {
	display: flex;
	align-items: baseline;
	gap: 0.75rem;
}
.detail-header h3 { margin-bottom: 0.5rem; }
.detail-header .delete {
	margin-left: auto;
	margin-bottom: 0;
	padding: 0.25rem 0.8rem;
	font-size: 0.85rem;
	width: auto;
	background: transparent;
	border-color: var(--pico-del-color);
	color: var(--pico-del-color);
}
.stage-strip {
	display: flex;
	flex-wrap: wrap;
	gap: 1.25rem;
	align-items: center;
	margin-bottom: 1rem;
}
.stage-group { display: flex; gap: 0.4rem; align-items: center; }
.stage-name {
	color: var(--pico-muted-color);
	font-size: 0.85rem;
	margin-right: 0.25rem;
}
.job-chip {
	display: inline-flex;
	align-items: center;
	gap: 0.35rem;
	padding: 0.35rem 0.7rem;
	border: 1px solid var(--pico-card-border-color);
	border-radius: 0.35rem;
	font-size: 0.85rem;
	cursor: pointer;
}
.job-chip.selected { border-color: var(--pico-primary); }
.log-header {
	display: flex;
	gap: 0.6rem;
	align-items: center;
	margin-bottom: 0.5rem;
}
.log-header .retry {
	margin-left: auto;
	margin-bottom: 0;
	padding: 0.25rem 0.8rem;
	font-size: 0.85rem;
	width: auto;
}
.log-pane pre {
	max-height: 70vh;
	overflow-y: auto;
	overflow-x: hidden;
	font-size: 0.8rem;
	white-space: pre-wrap;
	word-break: break-word;
}
`
