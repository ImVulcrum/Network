package main

import (
	"github.com/ImVulcrum/Network/network"
)

func main() {
	var n network.Network = network.New("My Network", false)

	n.AddKnot("A", []*network.Connection{}, []*network.Connection{})

	n.AddKnot("B", []*network.Connection{}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("A"), 5)})
	n.AddKnot("C", []*network.Connection{}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("A"), 2)})
	n.AddKnot("D", []*network.Connection{}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("A"), 1)})

	n.AddKnot("F", []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("B"), 1)}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("D"), 2)})
	n.AddKnot("E", []*network.Connection{}, []*network.Connection{n.ConnectionConstructor(n.GetKnotByContent("D"), 5)})

	n.Print("")

	table := n.Dijkstras(n.GetKnotByContent("A"))
	table.Print()
}
