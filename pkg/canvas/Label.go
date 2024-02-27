package canvas

// NewLabel ...
func NewLabel() LabelRepository {
	return &Label{}
}

/*
 * @Setters
 */

/* Elements */

/** DC **/

// SetBounds ...
func (label *Label) SetBounds() {
	label.Bounds = make([]Bounds, 1)
}

/*
 * @Getters
 */

/* Elements */

/** DC **/

// GetBounds ...
func (label Label) GetBounds() *Bounds {
	return &label.Bounds[0]
}
