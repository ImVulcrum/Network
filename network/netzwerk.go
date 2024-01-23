package network

type Network interface {
	String() string
	GiveCurrentKnot() *Knot
	GiveKnotCount() int
	//important functions
	AddKnot(inhalt string, outgoing_con []*Connection, incoming_con []*Connection) error
	ChangeNetworkName(new_name string)
	GiveNetworkName() string
	GetKnotByContent(content string) *Knot
	ConnectionConstructor(t *Knot, w int) *Connection
	DeleteKnotByContent(content string)
	MoveByWeight(type_of_movement int) error
	MoveByContent(content string) error
	SetCurrentKnot(knot *Knot)
	Print(tag string)

	Dijkstras(startknot *Knot) PriorityQueue
}
