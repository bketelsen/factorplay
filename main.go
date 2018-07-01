package main

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/bketelsen/factor/markup"
	"github.com/bketelsen/factorplay/div"
	"github.com/satori/go.uuid"
)

func main() {

	markup.Register(&div.Blue{})
	markup.Register(&div.Newdiv{})
	ctx := uuid.Must(uuid.NewV1())
	fmt.Println(ctx.String())
	c := &div.Blue{}
	c.Bar = "Brian"

	node, err := markup.Mount(c, ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, child := range node.Children {
		fmt.Println(child.Tag, child.ID)
	}
	fmt.Println(node.Markup())

	el := js.Global().Get("document").Call("getElementById", "thing")
	el.Set("innerHTML", node.Markup())
}
