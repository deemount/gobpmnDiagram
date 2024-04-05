package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

var (
	schemaBpmnDI = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDI  = "http://www.omg.org/spec/DD/20100524/DI"
)

// NewDefinitions ...
func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetBpmnDI ...
func (definitions *Definitions) SetBpmnDI() {
	definitions.BpmnDI = schemaBpmnDI
}

// SetOmgDI ...
func (definitions *Definitions) SetOmgDI() {
	definitions.OmgDI = schemaOMGDI
}

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram(num int) {
	definitions.Diagram = make(canvas.DIAGRAM_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** BPMNDI **/

// SetDiagram ...
func (definitions Definitions) GetDiagram(num int) *canvas.Diagram {
	return &definitions.Diagram[num]
}
