package canvas

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewShape ...
func NewShape() ShapeRepository {
	return &Shape{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (shape *Shape) SetID(typ string, suffix interface{}) {
	shape.ID = SetID(typ, suffix)
}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.Element = fmt.Sprintf("Activity_%s", suffix)
		break
	case "collaboration":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "event":
		shape.Element = fmt.Sprintf("Event_%s", suffix)
		break
	case "participant":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "startevent":
		shape.Element = fmt.Sprintf("StartEvent_%v", suffix)
		break
	case "id":
		shape.Element = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetIsHorizontal ...
func (shape *Shape) SetIsHorizontal(isHorizontal bool) {
	shape.IsHorizontal = isHorizontal
}

// SetIsMarkerVisible ...
func (shape *Shape) SetIsMarkerVisible(isMarkerVisible bool) {
	shape.IsMarkerVisible = isMarkerVisible
}

/* Elements */

/** DC **/

// SetBounds ...
func (shape *Shape) SetBounds() {
	shape.Bounds = make([]Bounds, 1)
}

/** BPMNDI **/

// SetLabel ...
func (shape *Shape) SetLabel() {
	shape.Label = make([]Label, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (shape Shape) GetID() impl.STR_PTR {
	return &shape.ID
}

// GetElement ...
func (shape Shape) GetElement() impl.STR_PTR {
	return &shape.Element
}

// GetIsHorizontal ...
func (shape Shape) GetIsHorizontal() impl.BOOL_PTR {
	return &shape.IsHorizontal
}

// GetIsMarkerVisible ...
func (shape Shape) GetIsMarkerVisible() impl.BOOL_PTR {
	return &shape.IsMarkerVisible
}

/* Elements */

/** DC **/

// GetBounds ...
func (shape Shape) GetBounds() *Bounds {
	return &shape.Bounds[0]
}

/** BPMNDI **/

// SetLabel ...
func (shape Shape) GetLabel() *Label {
	return &shape.Label[0]
}
