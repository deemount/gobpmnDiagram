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
	switch typ {
	case "diagram":
		//diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
		diagram.ID = fmt.Sprintf("BPMNDiagram_%v", suffix)
		break
	case "id":
		diagram.ID = fmt.Sprintf("%s", suffix)
		break
	}
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
