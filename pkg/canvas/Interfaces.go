package canvas

import impl "github.com/deemount/gobpmnTypes"

/*
 * @RBase
 */

// IFBaseCoordinates ...
type IFBaseCoordinates interface {
	SetCoordinates(x, y int)
	GetCoordinates() (impl.INT_PTR, impl.INT_PTR)
	SetX(x int)
	GetX() impl.INT_PTR
	SetY(y int)
	GetY() impl.INT_PTR
}

type CanvasBoundsElements interface {
	SetBounds()
	GetBounds() *Bounds
}

type CanvasLabelElements interface {
	SetLabel()
	GetLabel() *Label
}

/*
 * @Repositories
 */

// BoundsRepository ...
type BoundsRepository interface {
	IFBaseCoordinates
	SetSize(width, height int)
	GetSize() (impl.INT_PTR, impl.INT_PTR)
	SetWidth(width int)
	GetWidth() impl.INT_PTR
	SetHeight(height int)
	GetHeight() impl.INT_PTR
}

// WaypointRepository ...
type WaypointRepository interface {
	IFBaseCoordinates
}

// Diagram ...
type DiagramRepository interface {
	impl.IFBaseID
	SetPlane()
	GetPlane() *Plane
}

// Edge ...
type EdgeRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
	CanvasLabelElements

	SetWaypoint()
	SetWaypoints(num int)
	GetWaypoint(num int) *Waypoint
}

// Label ...
type LabelRepository interface {
	CanvasBoundsElements
}

// PlaneRepository ...
type PlaneRepository interface {
	impl.IFBaseID
	impl.IFBaseElement

	SetAttrProcessElement(suffix string)
	SetAttrCollaborationElement(suffix string)

	SetShape(num int)
	GetShape(num int) *Shape

	SetEdge(num int)
	GetEdge(num int) *Edge

	GetDescription() string
}

// ShapeRepository ...
type ShapeRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
	CanvasBoundsElements
	CanvasLabelElements

	SetIsHorizontal(isHorizontal bool)
	GetIsHorizontal() impl.BOOL_PTR

	SetIsMarkerVisible(isMarkerVisible bool)
	GetIsMarkerVisible() impl.BOOL_PTR
}
