package response

import "github.com/unrolled/render"

type renderer struct {
	*render.Render
}

var Render = renderer{}

func init() {
	// This would get options defined for the grouping.
	Render.Render = render.New(render.Options{})
}
