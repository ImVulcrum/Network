package network

import (
	"fmt"

	"../color"
)

type data struct {
	start                 []*Knoten
	knots                 []*Knoten
	aktKnoten             *Knoten
	name                  string
	allow_identical_names bool
}

type Knoten struct {
	kanten     []*Kante
	inhalt     string
	start_knot bool
}

type Kante struct {
	ziel    *Knoten
	gewicht int
}

func New(name string, allow_identical_names bool) *data {
	var n *data = new(data)
	n.name = name
	n.allow_identical_names = allow_identical_names
	return n
}

func (n *data) ChangeName(new_name string) {
	n.name = new_name
}

func (n *data) GiveName() string {
	return n.name
}

func (n *data) GibAktKnoten() *Knoten {
	return n.aktKnoten
}

func (n *data) GibAnzahlKnoten() int {
	return len(n.knots)
}

func (n *data) GetKnotByContent(content string) *Knoten {
	for i := 0; i < len(n.knots); i++ {
		if n.knots[i].inhalt == content {
			return n.knots[i]
		}
	}
	return nil
}

// func (k *Knoten) HatInhalt(inhalt string, anzahl, zaehler int) *Knoten {
// 	var nk *Knoten
// 	if anzahl == zaehler {
// 		return nk
// 	}
// 	if k.inhalt == inhalt {
// 		nk = k
// 	} else {
// 		for i := 0; i < len(k.kanten); i++ {
// 			k.kanten[i].ziel.HatInhalt(inhalt, anzahl, zaehler+1)
// 		}
// 	}
// 	return nk
// }

// func (n *data) Search(inhalt string) *Knoten {
// 	var k *Knoten
// 	var list []*Knoten
// 	for i := 0; i < len(n.start); i++ {
// 		list = append(list, n.start[i].HatInhalt(inhalt, len(n.knots), 0))
// 	}
// 	for i := 0; i < len(list); i++ {
// 		if list[i] != nil {
// 			k = list[i]
// 		}
// 	}
// 	return k
// }

// func (n *data) GehzuStart() {
// 	if len(n.start) != 0 {
// 		n.aktKnoten = n.start[0]
// 	}
// }

func (kn *Knoten) String() string {
	var erg string
	erg = erg + fmt.Sprintln(kn.inhalt)
	for i := 0; i < len(kn.kanten); i++ {
		erg = erg + "   --"
		erg = erg + fmt.Sprint(kn.kanten[i].gewicht, "-->")
		erg = erg + fmt.Sprintln(kn.kanten[i].ziel.inhalt)
	}
	return erg
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
	erg = erg + "aktKnoten:" + fmt.Sprintln(n.aktKnoten)
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
		if n.knots[i] == n.aktKnoten {
			color.Cyan(erg)
		} else {
			fmt.Print(erg)
		}
	}
	fmt.Print(fmt.Sprintln("\nAnzahl der Knoten:", len(n.knots)))
	if tag != "" {
		fmt.Println("Tag:", tag)
	}
	color.Red("-----" + n.name + "-----")
	fmt.Println("")
}

func (n *data) MoveByWeight(type_of_movement int) error {
	//-1 smallest
	//-2 highest

	if len(n.aktKnoten.kanten) == 0 {
		return fmt.Errorf("the current knot has no connections")
	}
	var index int = 0

	if type_of_movement >= 0 {
		var i int
		for i = 0; i < len(n.aktKnoten.kanten); i++ {
			if n.aktKnoten.kanten[i].gewicht == type_of_movement {
				n.aktKnoten = n.aktKnoten.kanten[i].ziel
				return nil
			}
		}
		if i == len(n.aktKnoten.kanten) {
			return fmt.Errorf("the current knot has no connection with he given weight")
		}
	} else if type_of_movement == -1 {
		var current_smallest int = n.aktKnoten.kanten[0].gewicht

		for i := 0; i < len(n.aktKnoten.kanten); i++ {
			if n.aktKnoten.kanten[i].gewicht < current_smallest {
				current_smallest = n.aktKnoten.kanten[i].gewicht
				index = i
			}
		}
	} else if type_of_movement == -2 {
		var current_highest int = n.aktKnoten.kanten[0].gewicht
		for i := 0; i < len(n.aktKnoten.kanten); i++ {
			if n.aktKnoten.kanten[i].gewicht > current_highest {
				current_highest = n.aktKnoten.kanten[i].gewicht
				index = i
			}
		}
	}
	n.aktKnoten = n.aktKnoten.kanten[index].ziel
	return nil
}

func (n *data) SetAktKnoten(knot *Knoten) {
	n.aktKnoten = knot
}

func (n *data) MoveToFirst() {
	if len(n.aktKnoten.kanten) != 0 {
		n.aktKnoten = n.aktKnoten.kanten[0].ziel
	}
}

func (kn *Knoten) MinNachbar() *Knoten {
	var nk *Knoten
	if len(kn.kanten) != 0 {
		var min int
		var index int
		min = kn.kanten[0].gewicht
		for i := 1; i < len(kn.kanten); i++ {
			if min > kn.kanten[i].gewicht {
				min = kn.kanten[i].gewicht
				index = i
			}
		}
		nk = kn.kanten[index].ziel
	}
	return nk
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

func (knot *Knoten) DeleteConnectionByDestination(delition_destination *Knoten) {
	var i int

	for i = 0; i < len(knot.kanten); i++ {
		if knot.kanten[i].ziel == delition_destination {
			break
		}
	}

	//delete function
	if i < len(knot.kanten) {
		knot.kanten[i] = knot.kanten[len(knot.kanten)-1]
		knot.kanten[len(knot.kanten)-1] = nil
		knot.kanten = knot.kanten[:len(knot.kanten)-1]
	}
}

func (n *data) NewConnection(t *Knoten, w int) *Kante {
	var con *Kante = new(Kante)
	con.ziel = t
	con.gewicht = w
	return con
}

func (n *data) AddKnot(inhalt string, outgoing_con []*Kante, incoming_con []*Kante) error {
	var knot *Knoten = new(Knoten)

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
		n.start = append(n.start, knot)
	} else {
		knot.start_knot = false
	}

	for i := 0; i < len(outgoing_con); i++ {
		knot.kanten = append(knot.kanten, outgoing_con[i])
		if outgoing_con[i].ziel.start_knot { //if one connection goes to a start knot this flag must be removed, cuz is is't a start knot anymore
			outgoing_con[i].ziel.start_knot = false
		}
	}

	for i := 0; i < len(incoming_con); i++ {
		var connection *Kante = new(Kante)
		connection.gewicht = incoming_con[i].gewicht
		connection.ziel = knot
		incoming_con[i].ziel.kanten = append(incoming_con[i].ziel.kanten, connection)
	}
	n.knots = append(n.knots, knot)

	return nil
}

// func (n *data) AddKnoten(inhalt string, qk *Knoten, gewicht int) {
// 	var k *Knoten = new(Knoten)
// 	k.inhalt = inhalt
// 	var nk *Kante = new(Kante)
// 	nk.ziel = k
// 	nk.gewicht = gewicht
// 	qk.kanten = append(qk.kanten, nk)
// 	// n.anzahl++
// }

// func (n *data) AddStartKnoten(inhalt string) {
// 	var k *Knoten = new(Knoten)
// 	k.inhalt = inhalt
// 	if len(n.knots) == 0 {
// 		n.start = append(n.start, k)
// 		// n.anzahl++
// 	}
// }

// func (n *data) AddFurtherStartknoten(inhalt string, zk *Knoten, gewicht int) {
// 	var k *Knoten = new(Knoten)
// 	k.inhalt = inhalt
// 	n.start = append(n.start, k)
// 	var nk *Kante = new(Kante)
// 	nk.ziel = zk
// 	nk.gewicht = gewicht
// 	k.kanten = append(k.kanten, nk)
// 	// n.anzahl++
// }

// func (n *data) AddKante(qk, zk *Knoten, gewicht int) {
// 	var nk *Kante = new(Kante)
// 	nk.ziel = zk
// 	nk.gewicht = gewicht
// 	qk.kanten = append(qk.kanten, nk)
// }
