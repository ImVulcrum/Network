package main

import (
	"fmt"

	"./network"
)

func main() {
	var n network.Network = network.New("My Network", false)

	n.AddKnot("knot1", []*network.Connection{}, []*network.Connection{})
	n.Print("")
	fmt.Println(n.AddKnot("knot1", []*network.Connection{}, []*network.Connection{}))
	n.Print("")
	n.AddKnot("knot2", []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("knot1"), 4)}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("knot1"), 2)})
	n.Print("")
	n.AddKnot("knot3", []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("knot1"), 4), n.ConnectionConstructor(n.GetKnotByContent("knot2"), 3)}, []*network.Connection{})

	n.Print("")

	n.SetCurrentKnot(n.GetKnotByContent("knot3"))

	n.Print("")

	n.MoveByWeight(4)

	n.Print("")

	n.MoveByWeight(2)

	n.Print("")
}
