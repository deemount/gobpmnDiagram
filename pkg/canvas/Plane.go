package canvas

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewPlane ...
func NewPlane() PlaneRepository {
	return &Plane{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMNDI **/

// SetID ...
func (plane *Plane) SetID(typ string, suffix interface{}) {
	switch typ {
	case "plane":
		plane.ID = fmt.Sprintf("BPMNPlane_%d", suffix)
	case "id":
		plane.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetElement ...
func (plane *Plane) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "process":
		plane.Element = fmt.Sprintf("Process_%s", suffix)
	case "collaboration":
		plane.Element = fmt.Sprintf("Collaboration_%s", suffix)
	case "id":
		plane.Element = fmt.Sprintf("%s", suffix)
	}
}

/*** Semantic ***/

// SetAttrProcessElement ...
func (plane *Plane) SetAttrProcessElement(suffix string) {
	plane.Element = fmt.Sprintf("Process_%s", suffix)
}

// SetAttrCollaborationElement ...
func (plane *Plane) SetAttrCollaborationElement(suffix string) {
	plane.Element = fmt.Sprintf("Collaboration_%s", suffix)
}

/* Elements */

/** BPMNDI **/

// SetShape ...
func (plane *Plane) SetShape(num int) {
	plane.Shape = make([]Shape, num)
}

// SetEdge ...
func (plane *Plane) SetEdge(num int) {
	plane.Edge = make([]Edge, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMNDI **/

// GetID ...
func (plane Plane) GetID() impl.STR_PTR {
	return &plane.ID
}

// GetElement ...
func (plane Plane) GetElement() impl.STR_PTR {
	return &plane.Element
}

/* Elements */

/** BPMNDI **/

// GetShape ...
func (plane Plane) GetShape(num int) *Shape {
	return &plane.Shape[num]
}

// GetEdge ...
func (plane Plane) GetEdge(num int) *Edge {
	return &plane.Edge[num]
}

/* Specification */

// Description ...
func (plane Plane) GetDescription() string {
	plane.Description = `
	A BPMNPlane specializes DI::Plane and redefines its model element
	reference to be of the (BPMN) BaseElement. A BPMNPlane can only reference
	a BaseElement of the types: Process, SubProcess, AdHocSubProcess, Transaction,
	Collaboration, Choreography or SubChoreography.

	BPMNPlane element is always owned by a BPMNDiagram and represents the root
	diagram element of that diagram. The plane represents a 2 dimensional surface
	with an origin at (0,0) along the x and y axes with increasing coordinates to
	the right and bottom. Only positive coordinates are allowed for diagram elements
	that are nested in a BPMNPlane. This means that the union of all the nested
	elements' bounds is deemed to be located at the plane's origin point.
	`
	return fmt.Sprintf("%s", plane.Description)
}
