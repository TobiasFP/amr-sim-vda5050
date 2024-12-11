package models

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

type ActionState struct {
	ActionID          string `json:"actionId"`
	ActionStatus      string `json:"actionStatus"`
	ActionType        string `json:"actionType"`
	ActionDescription string `json:"actionDescription"`
	ResultDescription string `json:"resultDescription"`
}
type BatteryState struct {
	BatteryCharge  float64 `json:"batteryCharge"`
	Charging       bool    `json:"charging"`
	BatteryVoltage float64 `json:"batteryVoltage"`
	BatteryHealth  float64 `json:"batteryHealth"`
	Reach          float64 `json:"reach"`
}
type Error struct {
	ErrorType        string        `json:"errorType"`
	ErrorLevel       string        `json:"errorLevel"`
	ErrorReferences  []interface{} `json:"errorReferences"`
	ErrorDescription string        `json:"errorDescription"`
	ErrorHint        string        `json:"errorHint"`
}

type Map struct {
	MapID          string `json:"mapId"`
	MapVersion     string `json:"mapVersion"`
	MapStatus      string `json:"mapStatus"`
	MapDescription string `json:"mapDescription"`
}
type SafetyState struct {
	EStop          string `json:"eStop"`
	FieldViolation bool   `json:"fieldViolation"`
}

type Info struct {
	InfoType       string `json:"infoType"`
	InfoLevel      string `json:"infoLevel"`
	InfoReferences []struct {
		ReferenceKey   string `json:"referenceKey"`
		ReferenceValue string `json:"referenceValue"`
	} `json:"infoReferences"`
	InfoDescription string `json:"infoDescription"`
}

type Load struct {
	LoadID               string `json:"loadId"`
	LoadType             string `json:"loadType"`
	LoadPosition         string `json:"loadPosition"`
	BoundingBoxReference struct {
		X     float64 `json:"x"`
		Y     float64 `json:"y"`
		Z     float64 `json:"z"`
		Theta float64 `json:"theta"`
	} `json:"boundingBoxReference"`
	LoadDimensions struct {
		Length float64 `json:"length"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	} `json:"loadDimensions"`
	Weight float64 `json:"weight"`
}

type State struct {
	HeaderID              int           `json:"headerId"`
	Timestamp             string        `json:"timestamp"`
	Version               string        `json:"version"`
	Manufacturer          string        `json:"manufacturer"`
	SerialNumber          string        `json:"serialNumber"`
	OrderID               string        `json:"orderId"`
	OrderUpdateID         int           `json:"orderUpdateId"`
	LastNodeID            string        `json:"lastNodeId"`
	LastNodeSequenceID    int           `json:"lastNodeSequenceId"`
	NodeStates            []NodeState   `json:"nodeStates"`
	EdgeStates            []EdgeState   `json:"edgeStates"`
	Driving               bool          `json:"driving"`
	ActionStates          []ActionState `json:"actionStates"`
	BatteryState          BatteryState  `json:"batteryState"`
	OperatingMode         string        `json:"operatingMode"`
	Errors                []Error       `json:"errors"`
	SafetyState           SafetyState   `json:"safetyState"`
	Maps                  []Map         `json:"maps"`
	ZoneSetID             string        `json:"zoneSetId"`
	Paused                bool          `json:"paused"`
	NewBaseRequest        bool          `json:"newBaseRequest"`
	DistanceSinceLastNode float64       `json:"distanceSinceLastNode"`
	AgvPosition           AgvPosition   `json:"agvPosition"`
	Velocity              Velocity      `json:"velocity"`
	Loads                 []Load        `json:"loads"`
	Information           []Info        `json:"information"`
}
