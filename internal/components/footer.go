package components

import (
	"strconv"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Footer struct {
	vecty.Core
}

func (footer *Footer) Render() vecty.ComponentOrHTML {
	now := time.Now().Year()
	return elem.Footer(
		elem.Div(
			vecty.Markup(
				vecty.Style("text-align", "center"),
				vecty.Style("padding", "20px"),
				vecty.Style("position", "fixed"),
				vecty.Style("width", "100%"),
				vecty.Style("bottom", "0"),
			),
			elem.Paragraph(
				vecty.Text("©️ "+strconv.Itoa(now)+" "),
				elem.Anchor(
					vecty.Markup(
						vecty.Property("href", "https://x.com/otakakot"),
					),
					vecty.Text("@otakakot"),
				),
				vecty.Text(" All rights reserved."),
			),
		),
	)
}
