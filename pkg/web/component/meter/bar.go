package meter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (m *Meter) bar() gomponents.Node {
	if m.gauge {
		share := 0.0

		if m.maximum > 0 {
			share = min(m.value/m.maximum, 1) * 100
		}

		return html.Div(
			html.Class(constant.MeterGaugeClass),
			html.Style(
				variables(
					constant.MeterShareVariable,
					fmt.Sprintf("%.1f%%", share),
					constant.MeterColorVariable,
					m.color,
					constant.MeterTrackVariable,
					m.track,
					constant.MeterWidthVariable,
					m.width,
					constant.MeterHeightVariable,
					m.height,
				),
			),
			html.Div(html.Class(constant.MeterGaugeFillClass)),
		)
	}

	style := variables(
		constant.ProgressColorVariable,
		m.color,
		constant.ProgressTrackVariable,
		m.track,
		constant.MeterWidthVariable,
		m.width,
		constant.MeterHeightVariable,
		m.height,
	)

	return html.Progress(
		html.Class(constant.MeterBarClass),
		html.Value(strconv.FormatFloat(m.value, 'f', -1, 64)),
		html.Max(strconv.FormatFloat(m.maximum, 'f', -1, 64)),
		gomponents.If(style != "", html.Style(style)),
	)
}
