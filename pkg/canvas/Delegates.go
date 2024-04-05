package canvas

import "fmt"

const startEvent = "startevent"
const ID = "id"
const participant = "participant"
const association = "association"
const activity = "activity"
const collaboration = "collaboration"
const event = "event"
const flow = "flow"
const process = "process"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case association:
		r = fmt.Sprintf("Association_%s_di", suffix)
	case activity:
		r = fmt.Sprintf("Activity_%s_di", suffix)
	case collaboration:
		r = fmt.Sprintf("Participant_%s_di", suffix)
	case participant:
		r = fmt.Sprintf("Participant_%s_di", suffix)
	case "di":
		r = fmt.Sprintf("%s_di", suffix)
	case event:
		r = fmt.Sprintf("Event_%s_di", suffix)
	case "diagram":
		r = fmt.Sprintf("BPMNDiagram_%v", suffix)
	case flow:
		r = fmt.Sprintf("Flow_%s_di", suffix)
	case "plane":
		r = fmt.Sprintf("BPMNPlane_%d", suffix)
	case startEvent:
		r = fmt.Sprintf("_BPMNShape_StartEvent_%v", suffix)
	case ID:
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}

func SetElement(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case activity:
		r = fmt.Sprintf("Activity_%s", suffix)
	case collaboration:
		r = fmt.Sprintf("Collaboration_%s", suffix)
	case event:
		r = fmt.Sprintf("Event_%s", suffix)
	case participant:
		r = fmt.Sprintf("Participant_%s", suffix)
	case process:
		r = fmt.Sprintf("Process_%s", suffix)
	case startEvent:
		r = fmt.Sprintf("StartEvent_%v", suffix)
	case ID:
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
