package network

import (
	"container/heap"
	"fmt"

	"color"
)

type data struct {
	knots                 []*Knot
	current_knot          *Knot
	name                  string
	allow_identical_names bool
}

type Knot struct {
	Connectionn []*Connection
	inhalt      string
	start_knot  bool
}

type Connection struct {
	ziel    *Knot
	gewicht int
}

type Item struct {
	cKnot         *Knot
	distance      int
	previous_knot *Knot
	visited       bool
}

// Priority Queue Implemenetation
type PriorityQueue []Item

// unnötig
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

//ende unnötig

// Less compares two elements in the priority queue based on their distances.
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].distance == -3 {
		return false
	} else if pq[j].distance == -3 {
		return true
	}

	return pq[i].distance < pq[j].distance
}

func (pq *PriorityQueue) RemoveAtIndex(index int) Item {
	old := *pq
	item := old[index]
	*pq = append(old[0:index], old[index+1:]...)
	heap.Init(pq) // Reinitialize the heap after removing an element
	return item
}

func (pq *PriorityQueue) ChangeElement(index int, new_distance int) {
	old_item := pq.RemoveAtIndex(index)
	old_item.distance = new_distance
	pq.AddElement(old_item)
}

func (pq PriorityQueue) GetIndexOfElementByKnot(knot *Knot) int {
	for i := 0; i < len(pq); i++ {
		if pq[i].cKnot == knot {
			return i
		}
	}

	return -1
}

func (pq PriorityQueue) Print() {
	for i := 0; i < len(pq); i++ {
		if pq[i].previous_knot == nil {
			fmt.Println(pq[i].cKnot.inhalt, pq[i].distance, "Keiner", pq[i].visited)
		} else {
			fmt.Println(pq[i].cKnot.inhalt, pq[i].distance, pq[i].previous_knot.inhalt, pq[i].visited)
		}
	}
}

func (pq PriorityQueue) GetIndexOfNearestNonVisitedItem() int {
	for i := 0; i < len(pq); i++ {
		if !pq[i].visited {
			return i
		}
	}

	return -1
}

func (pq *PriorityQueue) AddElement(item Item) {
	heap.Push(pq, item)
}

//End of Priority Queue Implementation

func New(name string, allow_identical_names bool) *data {
	var n *data = new(data)
	n.name = name
	n.allow_identical_names = allow_identical_names
	return n
}

func (n *data) ChangeNetworkName(new_name string) {
	n.name = new_name
}

func (n *data) GiveNetworkName() string {
	return n.name
}

func (n *data) GiveCurrentKnot() *Knot {
	return n.current_knot
}

func (n *data) GiveKnotCount() int {
	return len(n.knots)
}

func (n *data) GetKnotByContent(content string) *Knot {
	for i := 0; i < len(n.knots); i++ {
		if n.knots[i].inhalt == content {
			return n.knots[i]
		}
	}
	return nil
}

func (n *data) String() string {
	var erg string
	erg = fmt.Sprintln(n.name)
	for i := 0; i < len(n.knots); i++ {
		erg = erg + fmt.Sprint(i+1) + ". "
		if n.knots[i].start_knot {
			erg = erg + "Start Knot: "
		}
		erg = erg + fmt.Sprint(n.knots[i])
	}
	erg = erg + "current_knot:" + fmt.Sprintln(n.current_knot)
	erg = erg + fmt.Sprintln("Anzahl der Knoten:", len(n.knots))
	return erg
}

func (n *data) Print(tag string) {
	fmt.Println("")
	color.Red("-----" + n.name + "-----")
	for i := 0; i < len(n.knots); i++ {
		var erg = fmt.Sprint(i+1) + ". "
		if n.knots[i].start_knot {
			erg = erg + "StartKnot: "
		}
		erg = erg + fmt.Sprint(n.knots[i])
		if n.knots[i] == n.current_knot {
			color.Cyan(erg)
		} else {
			fmt.Print(erg)
		}
	}
	fmt.Print(fmt.Sprintln("\nAnzahl der Knot:", len(n.knots)))
	if tag != "" {
		fmt.Println("Tag:", tag)
	}
	color.Red("-----" + n.name + "-----")
	fmt.Println("")
}

func (n *data) MoveByContent(content string) error {
	if len(n.current_knot.Connectionn) == 0 {
		return fmt.Errorf("the current knot has no connections")
	}

	var i int
	for i = 0; i < len(n.current_knot.Connectionn); i++ {
		if n.current_knot.Connectionn[i].ziel.inhalt == content {
			n.current_knot = n.current_knot.Connectionn[i].ziel
			return nil
		}
	}

	if i == len(n.current_knot.Connectionn) {
		return fmt.Errorf("the current knot has no connection to the given knot")
	}
	return nil
}

func (n *data) MoveByWeight(type_of_movement int) error {
	//-1 smallest
	//-2 highest

	if len(n.current_knot.Connectionn) == 0 {
		return fmt.Errorf("the current knot has no connections")
	}
	var index int = 0

	if type_of_movement >= 0 {
		var i int
		for i = 0; i < len(n.current_knot.Connectionn); i++ {
			if n.current_knot.Connectionn[i].gewicht == type_of_movement {
				n.current_knot = n.current_knot.Connectionn[i].ziel
				return nil
			}
		}
		if i == len(n.current_knot.Connectionn) {
			return fmt.Errorf("the current knot has no connection with he given weight")
		}
	} else if type_of_movement == -1 {
		var current_smallest int = n.current_knot.Connectionn[0].gewicht

		for i := 0; i < len(n.current_knot.Connectionn); i++ {
			if n.current_knot.Connectionn[i].gewicht < current_smallest {
				current_smallest = n.current_knot.Connectionn[i].gewicht
				index = i
			}
		}
	} else if type_of_movement == -2 {
		var current_highest int = n.current_knot.Connectionn[0].gewicht
		for i := 0; i < len(n.current_knot.Connectionn); i++ {
			if n.current_knot.Connectionn[i].gewicht > current_highest {
				current_highest = n.current_knot.Connectionn[i].gewicht
				index = i
			}
		}
	}
	n.current_knot = n.current_knot.Connectionn[index].ziel
	return nil
}

func (n *data) SetCurrentKnot(knot *Knot) {
	n.current_knot = knot
}

func (n *data) DeleteKnotByContent(content string) {
	var i int

	//get index of not to be deleted in knots slice
	for i = 0; i < len(n.knots); i++ {
		if n.knots[i].inhalt == content {
			break
		}
	}

	//delete all connections going in the not to be deleted
	for k := 0; k < len(n.knots); k++ {
		n.knots[k].DeleteConnectionByDestination(n.knots[i])
	}

	//delete from knots slice
	if i < len(n.knots) {
		n.knots[i] = n.knots[len(n.knots)-1]
		n.knots[len(n.knots)-1] = nil
		n.knots = n.knots[:len(n.knots)-1]
	}
}

func (n *data) ConnectionConstructor(t *Knot, w int) *Connection {
	var con *Connection = new(Connection)
	con.ziel = t
	con.gewicht = w
	return con
}

func (n *data) AddKnot(inhalt string, outgoing_con []*Connection, incoming_con []*Connection) error {
	var knot *Knot = new(Knot)

	if !n.allow_identical_names {
		for i := 0; i < len(n.knots); i++ {
			if n.knots[i].inhalt == inhalt {
				return fmt.Errorf("a knot with the same name already exist, please consider another name")
			}
		}
	}

	knot.inhalt = inhalt

	if len(incoming_con) == 0 && (len(outgoing_con) != 0 || len(n.knots) == 0) { //if the knot has no incoming knots but outcoming knots or also no outcoming knots if it is the first knot
		knot.start_knot = true
	} else {
		knot.start_knot = false
	}

	for i := 0; i < len(outgoing_con); i++ {
		knot.Connectionn = append(knot.Connectionn, outgoing_con[i])
		if outgoing_con[i].ziel.start_knot { //if one connection goes to a start knot this flag must be removed, cuz is is't a start knot anymore
			outgoing_con[i].ziel.start_knot = false
		}
	}

	for i := 0; i < len(incoming_con); i++ {
		var connection *Connection = new(Connection)
		connection.gewicht = incoming_con[i].gewicht
		connection.ziel = knot
		incoming_con[i].ziel.Connectionn = append(incoming_con[i].ziel.Connectionn, connection)
	}
	n.knots = append(n.knots, knot)

	return nil
}

//knot functions

func (knot *Knot) String() string {
	var erg string
	erg = erg + fmt.Sprintln(knot.inhalt)
	for i := 0; i < len(knot.Connectionn); i++ {
		erg = erg + "   --"
		erg = erg + fmt.Sprint(knot.Connectionn[i].gewicht, "-->")
		erg = erg + fmt.Sprintln(knot.Connectionn[i].ziel.inhalt)
	}
	return erg
}

func (knot *Knot) DeleteConnectionByDestination(delition_destination *Knot) {
	var i int

	for i = 0; i < len(knot.Connectionn); i++ {
		if knot.Connectionn[i].ziel == delition_destination {
			break
		}
	}

	//delete function
	if i < len(knot.Connectionn) {
		knot.Connectionn[i] = knot.Connectionn[len(knot.Connectionn)-1]
		knot.Connectionn[len(knot.Connectionn)-1] = nil
		knot.Connectionn = knot.Connectionn[:len(knot.Connectionn)-1]
	}
}

func (n *data) Dijkstras(startknot *Knot) PriorityQueue {
	n.SetCurrentKnot(startknot)

	table := make(PriorityQueue, 0)

	for i := 0; i < len(n.knots); i++ {
		if n.knots[i] == startknot {
			table.AddElement(Item{cKnot: startknot, distance: 0, previous_knot: nil, visited: true})
		} else {
			table.AddElement(Item{cKnot: n.knots[i], distance: -3, previous_knot: nil, visited: false})
		}
	}

	for {
		connections := n.current_knot.Connectionn
		for i := 0; i < len(connections); i++ {
			//überprüfen ob kanten gewicht + gewicht von wo wir sind kleiner als das was schon in der tabelle steht

			current_distant_to_knot := table[table.GetIndexOfElementByKnot(n.GiveCurrentKnot())].distance + connections[i].gewicht

			if current_distant_to_knot < table[table.GetIndexOfElementByKnot(connections[i].ziel)].distance || table[table.GetIndexOfElementByKnot(connections[i].ziel)].distance == -3 {
				table.ChangeElement(table.GetIndexOfElementByKnot(connections[i].ziel), current_distant_to_knot)
				table[table.GetIndexOfElementByKnot(connections[i].ziel)].previous_knot = n.GiveCurrentKnot()
			}
		}

		index_of_nearest_item := table.GetIndexOfNearestNonVisitedItem()

		if index_of_nearest_item == -1 {
			break
		}

		table[index_of_nearest_item].visited = true
		//table[index_of_nearest_item].previous_knot = n.GiveCurrentKnot()
		n.SetCurrentKnot(table[index_of_nearest_item].cKnot)
	}

	return table
}
