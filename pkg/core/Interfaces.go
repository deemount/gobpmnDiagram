package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

/*
 * @Repositories
 */

// DefinitionsRepository ...
type DefinitionsRepository interface {
	SetDiagram(num int)
	GetDiagram(num int) *canvas.Diagram
	SetBpmnDI()
	SetOmgDI()
}
