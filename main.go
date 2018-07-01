package main

import (
	"log"

	"github.com/bketelsen/factor/markup"
	"github.com/bketelsen/factorplay/div"
)

func main() {

	c := &div.Newdiv{}
	c.Bar = "Bar"
	_, err := markup.MountBody(c)
	if err != nil {
		log.Fatal(err)
	}
}
