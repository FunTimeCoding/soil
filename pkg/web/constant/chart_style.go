package constant

const ChartStyle = `
.chart svg { width: 100%; height: auto; }
.chart-grid { stroke: var(--chart-grid, rgba(128,128,128,0.2)); stroke-width: 1; }
.chart-ceiling { stroke: var(--chart-ceiling, rgba(128,128,128,0.55)); stroke-width: 1; stroke-dasharray: 2 3; }
.chart-guide { stroke: var(--chart-guide, rgba(128,128,128,0.45)); stroke-width: 1; stroke-dasharray: 6 4; }
.chart-now { stroke: var(--chart-now, rgba(220,80,80,0.7)); stroke-width: 1; }
.chart-line { fill: none; stroke-width: 2; stroke-linejoin: round; }
.chart-projection { stroke-width: 1.5; stroke-dasharray: 3 4; opacity: 0.8; }
.chart-label { fill: var(--chart-label, rgba(128,128,128,0.9)); font-size: 10px; stroke: none; }
.chart-label-value { text-anchor: end; }
.chart-label-day { text-anchor: middle; }
.chart-projection-label { font-size: 11px; }
.chart-legend { display: flex; gap: 1.2em; font-size: 0.85em; margin-top: 0.3em; }
.chart-legend-item { display: inline-flex; align-items: center; gap: 0.4em; }
.chart-legend-dot { width: 0.7em; height: 0.7em; border-radius: 50%; display: inline-block; background: currentColor; }
.chart-series-a { stroke: var(--chart-series-a, #4f8edc); color: var(--chart-series-a, #4f8edc); }
.chart-series-b { stroke: var(--chart-series-b, #d9962e); color: var(--chart-series-b, #d9962e); }
.chart-projection-label.chart-series-a { fill: var(--chart-series-a, #4f8edc); }
.chart-projection-label.chart-series-b { fill: var(--chart-series-b, #d9962e); }
`
