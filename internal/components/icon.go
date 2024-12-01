package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Icon struct {
	vecty.Core
	HRef string
}

func (icon *Icon) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Text("X ( Twitter )"),
		vecty.Markup(
			vecty.Property("href", icon.HRef),
		),
	)
}
