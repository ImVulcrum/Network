package main

import (
	"fmt"

	"./network"
)

func main() {
	var n network.Netzwerk = network.New("My Network", false)

	n.AddKnot("knot1", []*network.Kante{}, []*network.Kante{})
	n.Print("")
	fmt.Println(n.AddKnot("knot1", []*network.Kante{}, []*network.Kante{}))
	n.Print("")
	n.AddKnot("knot2", []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 4)}, []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 2)})
	n.Print("")
	n.AddKnot("knot3", []*network.Kante{n.NewConnection(n.GetKnotByContent("knot1"), 4), n.NewConnection(n.GetKnotByContent("knot2"), 3)}, []*network.Kante{})

	n.Print("")

	n.SetAktKnoten(n.GetKnotByContent("knot3"))

	n.Print("")

	n.MoveByWeight(4)

	n.Print("")

	n.MoveByWeight(2)

	n.Print("")
}
