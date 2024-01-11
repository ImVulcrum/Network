package network

import (
	"fmt"

	"../color"
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
	erg = erg + fmt.Sprintln("Anzahl der Knot:", len(n.knots))
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
