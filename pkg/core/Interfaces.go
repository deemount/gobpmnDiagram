package core

import "github.com/deemount/gobpmnDiagram/pkg/canvas"

/*
 * @Base
 */

// DefinitionsElements ...
type DefinitionsElements interface {
	SetDiagram(num int)
	GetDiagram(num int) canvas.DIAGRAM_PTR
}

/*
 * @Repositories
 */

// DefinitionsRepository ...
type DefinitionsRepository interface {
	DefinitionsElements
	SetBpmnDI()
	SetOmgDI()
}
