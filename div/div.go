package div

import (
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

func (d *Blue) OnInputChange(v string) {
	d.Bar = v
	d.Render()
}
func (d *Blue) Render() string {
	return box.String("blue.html")
}

func init() {
	box = packr.NewBox("./templates")
}
