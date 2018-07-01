package div

import (
	"github.com/bketelsen/factor/markup"
	"github.com/gobuffalo/packr"
)

var box packr.Box

type Newdiv struct {
	Bar string
}

func (d *Newdiv) Render() string {
	return box.String("newdiv.html")

}

type Blue struct {
	Bar string
}

func (d *Blue) Render() string {
	return box.String("blue.html")
}

type Body struct {
	Bar string
}

func (d *Body) Render() string {
	return box.String("body.html")
}
func init() {
	markup.Register(&Blue{})
	markup.Register(&Newdiv{})
	markup.Register(&Body{})
	box = packr.NewBox("./templates")
}
