package canvas

import impl "github.com/deemount/gobpmnTypes"

// NewWaypoint ...
func NewWaypoint() WaypointRepository {
	return &Waypoint{}
}

/*
 * @Setters
 */

/* Attributes */

// SetX ...
func (wp *Waypoint) SetX(x int) {
	wp.X = x
}

// SetY ...
func (wp *Waypoint) SetY(y int) {
	wp.Y = y
}

// SetCoordinates ...
func (wp *Waypoint) SetCoordinates(x, y int) {
	wp.X = x
	wp.Y = y
}

/*
 * @Getters
 */

/* Attributes */

// GetX ...
func (wp Waypoint) GetX() impl.INT_PTR {
	return &wp.X
}

// GetY ...
func (wp Waypoint) GetY() impl.INT_PTR {
	return &wp.Y
}

// GetCoordinates ...
func (wp Waypoint) GetCoordinates() (impl.INT_PTR, impl.INT_PTR) {
	return &wp.X, &wp.Y
}
