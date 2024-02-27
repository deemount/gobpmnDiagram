package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

// DefinitionsBaseElements ...
type DefinitionsBaseElements struct {
	Diagram []canvas.Diagram `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
}

// TDefinitionsBaseElements ...
type TDefinitionsBaseElements struct {
	Diagram []canvas.TDiagram `xml:"BPMNDiagram,omitempty" json:"diagram"`
}

/*
 * @Elementary
 */

// Definitions represents the root element
type Definitions struct {
	BpmnDI string `xml:"xmlns:bpmndi,attr" json:"-"`
	OmgDI  string `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
	DefinitionsBaseElements
}

// TDefinitions ...
type TDefinitions struct {
	TDefinitionsBaseElements
}
