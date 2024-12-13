package models

type Velocity struct {
	Vx    float64 `json:"vx"`
	Vy    float64 `json:"vy"`
	Omega float64 `json:"omega"`
}

type AgvPosition struct {
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
	Theta               float64 `json:"theta"`
	MapID               string  `json:"mapId"`
	PositionInitialized bool    `json:"positionInitialized"`
	MapDescription      string  `json:"mapDescription"`
	LocalizationScore   float64 `json:"localizationScore"`
	DeviationRange      float64 `json:"deviationRange"`
}

type NodeState struct {
	NodeID          string "json:\"nodeId\""
	SequenceID      int    "json:\"sequenceId\""
	Released        bool   "json:\"released\""
	NodeDescription string "json:\"nodeDescription\""
	NodePosition    struct {
		X     float64 "json:\"x\""
		Y     float64 "json:\"y\""
		MapID string  "json:\"mapId\""
		Theta float64 "json:\"theta\""
	} "json:\"nodePosition\""
}

type EdgeState struct {
	EdgeID          string `json:"edgeId"`
	SequenceID      int    `json:"sequenceId"`
	Released        bool   `json:"released"`
	EdgeDescription string `json:"edgeDescription"`
	Trajectory      struct {
		Degree        int           `json:"degree"`
		KnotVector    []interface{} `json:"knotVector"`
		ControlPoints []struct {
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Weight float64 `json:"weight"`
		} `json:"controlPoints"`
	} `json:"trajectory"`
}
