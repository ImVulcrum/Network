package main

import (
	"fmt"

	"./network"
)

func main() {
	var n network.Netzwerk = network.New("My Network", false)

	n.AddKnot("knot1", []*network.Kante{}, []*network.Kante{})
	fmt.Println(n)
	fmt.Println(n.AddKnot("knot1", []*network.Kante{}, []*network.Kante{}))
	fmt.Println(n)
	n.AddKnot("knot2", []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 4)}, []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 4)})

	n.AddKnot("knot3", []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 4), n.NewConnection(n.GetKnotByContent("knot2"), 4)}, []*network.Kante{})

	fmt.Println(n)

	n.DeleteKnotByContent("knot2")

	fmt.Println(n)
}
