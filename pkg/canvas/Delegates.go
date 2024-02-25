package canvas

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "association":
		r = fmt.Sprintf("Association_%s_di", suffix)
		break
	case "activity":
		r = fmt.Sprintf("Activity_%s_di", suffix)
		break
	case "collaboration":
		r = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "di":
		r = fmt.Sprintf("%s_di", suffix)
		break
	case "event":
		r = fmt.Sprintf("Event_%s_di", suffix)
		break
	case "diagram":
		//diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
		r = fmt.Sprintf("BPMNDiagram_%v", suffix)
		break
	case "flow":
		r = fmt.Sprintf("Flow_%s_di", suffix)
		break
	case "participant":
		r = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "plane":
		//r = "BPMNPlane_" + strconv.FormatInt(num, 16)
		r = fmt.Sprintf("BPMNPlane_%d", suffix)
		break
	case "startevent":
		r = fmt.Sprintf("_BPMNShape_StartEvent_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}

func SetElement(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "activity":
		r = fmt.Sprintf("Activity_%s", suffix)
		break
	case "collaboration":
		r = fmt.Sprintf("Participant_%s", suffix)
		break
	case "event":
		r = fmt.Sprintf("Event_%s", suffix)
		break
	case "participant":
		r = fmt.Sprintf("Participant_%s", suffix)
		break
	case "startevent":
		r = fmt.Sprintf("StartEvent_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
