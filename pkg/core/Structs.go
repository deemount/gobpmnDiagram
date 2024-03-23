package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

/*
 * @Elementary
 */

// Definitions represents the root element
type Definitions struct {
	BpmnDI  string           `xml:"xmlns:bpmndi,attr" json:"-"`
	OmgDI   string           `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
	Diagram []canvas.Diagram `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
}

// TDefinitions ...
type TDefinitions struct {
	Diagram []canvas.TDiagram `xml:"BPMNDiagram,omitempty" json:"diagram"`
}
