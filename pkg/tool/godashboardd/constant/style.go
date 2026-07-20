package constant

const InlineStyle = `
.board-grid { display: grid; gap: 0.75rem; align-items: start; }
.board-tail { margin-top: 2.5rem; }
.board-column { display: flex; flex-direction: column; gap: 0.75rem; }
.board-section { margin: 0; padding: 0.75rem 1rem; }
.board-section h4 { margin-bottom: 0.25rem; font-size: 0.85rem; text-transform: uppercase; letter-spacing: 0.05em; color: var(--pico-muted-color); }
.board-entry-link { display: flex; align-items: center; gap: 0.6rem; padding: 0.3rem 0; }
.board-icon { width: 1.4rem; height: 1.4rem; object-fit: contain; flex-shrink: 0; }
.board-rows { display: grid; grid-template-columns: repeat(auto-fit, minmax(5rem, 1fr)); gap: 0.15rem 0.75rem; margin: 0 0 0.3rem 2rem; font-size: 0.8rem; color: var(--pico-muted-color); }
.board-row-value { font-weight: 600; color: var(--pico-color); }
`
