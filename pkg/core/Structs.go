package core

import (
	"encoding/xml"

	"github.com/deemount/gobpmnDiagram/pkg/canvas"
)

/*
 * @Elementary
 */

// Definitions represents the root element
type (
	Definitions struct {
		XMLName xml.Name         `xml:"bpmn:definitions" json:"-"`
		BpmnDI  string           `xml:"xmlns:bpmndi,attr" json:"-"`
		OmgDI   string           `xml:"xmlns:omgdi,attr,omitempty" json:"-"`
		Diagram []canvas.Diagram `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram"`
	}

	// TDefinitions ...
	TDefinitions struct {
		Diagram []canvas.TDiagram `xml:"BPMNDiagram,omitempty" json:"diagram"`
	}
)
