package network

type Netzwerk interface {
	// AddStartKnoten(inhalt string)
	// Search(inhalt string) *Knoten
	// AddKnoten(inhalt string, qk *Knoten, gewicht int)
	// AddFurtherStartknoten(inhalt string, zk *Knoten, gewicht int)
	// AddKante(qk, zk *Knoten, gewicht int)
	// GehzuStart()
	String() string
	GibAktKnoten() *Knoten
	GibAnzahlKnoten() int
	MoveToFirst()
	//important functions
	AddKnot(inhalt string, outgoing_con []*Kante, incoming_con []*Kante) error
	ChangeName(new_name string)
	GiveName() string
	GetKnotByContent(content string) *Knoten
	NewConnection(t *Knoten, w int) *Kante
	DeleteKnotByContent(content string)
}
