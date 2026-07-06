package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func raidDetailTable(rows []store.RaidPlayerRow) gomponents.Node {
	if len(rows) == 0 {
		return html.P(html.Em(gomponents.Text("No player stats for this raid.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Player")),
				html.Th(gomponents.Text("Profession")),
				html.Th(gomponents.Text("Fights")),
				html.Th(gomponents.Text("Dmg/s")),
				html.Th(gomponents.Text("Heal/s")),
				html.Th(gomponents.Text("Clean/s")),
				html.Th(gomponents.Text("Strip/s")),
				html.Th(gomponents.Text("Barrier/s")),
				html.Th(gomponents.Text("Downs/s")),
				html.Th(gomponents.Text("Death/m")),
				html.Th(gomponents.Text("Tag Dist")),
			),
		),
		html.TBody(
			gomponents.Map(
				rows,
				func(r store.RaidPlayerRow) gomponents.Node {
					seconds := float64(r.ActiveTimeMS) / 1000
					minutes := seconds / 60
					icon := fmt.Sprintf(
						"/static/icons/%s.png",
						strings.ToLower(r.Profession),
					)

					return html.Tr(
						html.Td(
							html.A(
								html.Href(
									fmt.Sprintf(
										"/players/%s",
										r.Account,
									),
								),
								gomponents.Text(r.Name),
							),
						),
						html.Td(
							html.Img(
								html.Src(icon),
								html.Width("20"),
								html.Height("20"),
							),
							gomponents.Textf(" %s", r.Profession),
						),
						html.Td(gomponents.Textf("%d", r.Fights)),
						html.Td(
							gomponents.Text(
								perSecond(
									r.Damage,
									seconds,
									1,
								),
							),
						),
						html.Td(
							gomponents.Text(
								perSecond(
									r.Healing,
									seconds,
									1,
								),
							),
						),
						html.Td(
							gomponents.Text(
								perSecond(r.ConditionCleanses, seconds, 1),
							),
						),
						html.Td(
							gomponents.Text(
								perSecond(
									r.BoonStrips,
									seconds,
									2,
								),
							),
						),
						html.Td(
							gomponents.Text(
								perSecond(
									r.Barrier,
									seconds,
									1,
								),
							),
						),
						html.Td(
							gomponents.Text(
								perSecond(
									r.Downs,
									seconds,
									1,
								),
							),
						),
						html.Td(
							gomponents.Text(
								perMinute(
									r.DeadCount,
									minutes,
								),
							),
						),
						html.Td(gomponents.Textf("%.0f", r.DistToCom)),
					)
				},
			),
		),
	)
}
