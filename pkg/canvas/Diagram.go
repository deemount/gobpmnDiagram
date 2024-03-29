package canvas

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewDiagram ...
func NewDiagram() DiagramRepository {
	return &Diagram{}
}

/*
 * @Setters
 */

/* Attributes */

// SetID ...
// Notice: old fashion style SetID(typ string, num int64); look also below
func (diagram *Diagram) SetID(typ string, suffix interface{}) {
	diagram.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMNDI **/

// SetPlane ...
func (diagram *Diagram) SetPlane() {
	diagram.Plane = make([]Plane, 1)
}

/*
 * @Getters
 */

/* Attributes */

// GetID ...
func (diagram Diagram) GetID() impl.STR_PTR {
	return &diagram.ID
}

/* Elements */

/** BPMNDI **/

// GetPlane ...
func (diagram Diagram) GetPlane() *Plane {
	return &diagram.Plane[0]
}

/* Specification */

// Description ...
func (diagram Diagram) GetDescription() string {
	diagram.Description = `
	BPMNDiagram represents a depiction of all or part of a BPMN model. It
	specializes DI::Diagram and redefines the root element (the top most diagram element)
	to be of type BPMNPlane. A BPMN diagram can also own a collection of BPMNStyle elements
	that are referenced by BPMNLabel elements in the diagram. These style elements represent
	the unique appearance styles used in the diagram.
	`
	return fmt.Sprintf("%s", diagram.Description)
}
