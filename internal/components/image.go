package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Image struct {
	vecty.Core
	Src    string
	Widch  string
	Height string
}

func (image *Image) Render() vecty.ComponentOrHTML {
	return elem.Image(
		vecty.Markup(
			vecty.Style("width", image.Widch),
			vecty.Style("height", image.Height),
			vecty.Property("src", image.Src),
		),
	)
}
