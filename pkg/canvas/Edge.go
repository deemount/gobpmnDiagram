package canvas

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewEdge ...
func NewEdge() EdgeRepository {
	return &Edge{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (edge *Edge) SetID(typ string, suffix interface{}) {
	edge.ID = SetID(typ, suffix)
}

// SetElement ...
func (edge *Edge) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "association":
		edge.Element = fmt.Sprintf("Association_%s", suffix)
	case "flow":
		edge.Element = fmt.Sprintf("Flow_%s", suffix)
	}
}

/* Elements */

/** BPMNDI **/

// SetWaypoint ...
func (edge *Edge) SetWaypoint() {
	edge.Waypoint = make([]Waypoint, 2)
}

// SetWaypoints ...
func (edge *Edge) SetWaypoints(num int) {
	edge.Waypoint = make([]Waypoint, num)
}

// SetLabel ...
func (edge *Edge) SetLabel() {
	edge.Label = make([]Label, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (edge Edge) GetID() impl.STR_PTR {
	return &edge.ID
}

// GetElement ...
func (edge Edge) GetElement() impl.STR_PTR {
	return &edge.Element
}

/* Elements */

/** BPMNDI **/

// GetWaypoint ...
func (edge Edge) GetWaypoint(num int) *Waypoint {
	return &edge.Waypoint[num]
}

// GetLabel ...
func (edge Edge) GetLabel() *Label {
	return &edge.Label[0]
}
