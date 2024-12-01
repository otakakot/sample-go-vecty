package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/otakakot/sample-go-vecty/internal/components"
	"github.com/otakakot/sample-go-vecty/internal/meta"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			elem.Heading1(vecty.Text("Hello Vecty!")),
			elem.Heading2(vecty.Text("Hello Vecty!")),
			elem.Heading3(vecty.Text("Hello Vecty!")),
			elem.Heading1(
				vecty.Markup(
					vecty.Style("color", "red"),
					vecty.Style("text-align", "center"),
				),
				vecty.Text("Hello Vecty!"),
			),
			&components.Icon{
				HRef: "https://x.com/otakakot",
			},
			&components.Image{
				Src: "https://go.dev/doc/gopher/frontpage.png",
			},
			&components.Image{
				Src:    "/public/profile.png",
				Widch:  "100px",
				Height: "100px",
			},
			vecty.Text("Hello Vecty!"),
		),
		&components.Footer{},
	)
}

func main() {
	vecty.SetTitle("Vecty Page!")
	meta.Set("description", "", "A simple Vecty page.")
	vecty.RenderBody(&PageView{})
}
