package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

var (
	schemaBpmnDI = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDI  = "http://www.omg.org/spec/DD/20100524/DI"
)

/*
 * Default Setters
 */

/* Elements */

/** BPMNDI **/

// SetDiagram ...
func (definitions *Definitions) SetDiagram(num int) {
	definitions.Diagram = make(canvas.DIAGRAM_SLC, num)
}

/*
 * Default Getters
 */

/* Elements */

/** BPMNDI **/

// SetDiagram ...
func (definitions Definitions) GetDiagram(num int) canvas.DIAGRAM_PTR {
	return &definitions.Diagram[num]
}
