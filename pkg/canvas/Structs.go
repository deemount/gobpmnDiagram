package canvas

/*
 * Elementary
 */

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty" json:"x,omitempty"`
	Y      int `xml:"y,attr,omitempty" json:"y,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"x,omitempty"`
	Y int `xml:"y,attr" json:"y,omitempty"`
}

// Diagram ...
type Diagram struct {
	ID          string    `xml:"id,attr" json:"id,omitempty"`
	Description string    `xml:"-" json:"-"`
	Plane       PLANE_SLC `xml:"bpmndi:BPMNPlane,omitempty" json:"plane,omitempty"`
}

// TDiagram ...
type TDiagram struct {
	ID          string     `xml:"id,attr" json:"id,omitempty"`
	Description string     `xml:"-" json:"-"`
	Plane       TPLANE_SLC `xml:"BPMNPlane,omitempty" json:"plane,omitempty"`
}

// Edge ...
type Edge struct {
	ID       string       `xml:"id,attr" json:"-"`
	Element  string       `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint WAYPOINT_SLC `xml:"di:waypoint" json:"waypoint,omitempty"`
	Label    LABEL_SLC    `xml:"bpmndi:BPMNLabel,omitempty" json:"label,omitempty"`
}

// TEdge ...
type TEdge struct {
	ID       string       `xml:"id,attr" json:"-"`
	Element  string       `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint WAYPOINT_SLC `xml:"waypoint" json:"waypoint,omitempty"`
	Label    LABEL_SLC    `xml:"BPMNLabel,omitempty" json:"label,omitempty"`
}

// Label ...
type Label struct {
	Bounds BOUNDS_SLC `xml:"dc:Bounds,omitempty" json:"label,omitempty"`
}

// TLabel ...
type TLabel struct {
	Bounds BOUNDS_SLC `xml:"Bounds,omitempty" json:"bounds,omitempty"`
}

// Plane ...
type Plane struct {
	ID          string    `xml:"id,attr" json:"id,omitempty"`
	Element     string    `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Description string    `xml:"-" json:"-"`
	Shape       SHAPE_SLC `xml:"bpmndi:BPMNShape" json:"shape,omitempty"`
	Edge        EDGE_SLC  `xml:"bpmndi:BPMNEdge" json:"edge,omitempty"`
}

// TPlane ...
type TPlane struct {
	ID          string     `xml:"id,attr" json:"id,omitempty"`
	Element     string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Description string     `xml:"-" json:"-"`
	Shape       TSHAPE_SLC `xml:"BPMNShape" json:"shape,omitempty"`
	Edge        TEDGE_SLC  `xml:"BPMNEdge" json:"edge,omitempty"`
}

// Shape ...
type Shape struct {
	ID              string     `xml:"id,attr" json:"id,omitempty"`
	Element         string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	IsHorizontal    bool       `xml:"isHorizontal,attr,omitempty" json:"isHorizontal,omitempty"`
	IsMarkerVisible bool       `xml:"isMarkerVisible,attr,omitempty" json:"isMarkerVisible,omitempty"`
	Bounds          BOUNDS_SLC `xml:"dc:Bounds" json:"bounds,omitempty"`
	Label           LABEL_SLC  `xml:"bpmndi:BPMNLabel" json:"label,omitempty"`
}

// TShape ...
type TShape struct {
	ID              string     `xml:"id,attr" json:"id,omitempty"`
	Element         string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	IsHorizontal    bool       `xml:"isHorizontal,attr,omitempty" json:"isHorizontal,omitemptyy"`
	IsMarkerVisible bool       `xml:"isMarkerVisible,attr,omitempty" json:"isMarkerVisible,omitempty"`
	Bounds          BOUNDS_SLC `xml:"Bounds" json:"bounds,omitempty"`
	Label           TLABEL_SLC `xml:"BPMNLabel" json:"label,omitempty"`
}
