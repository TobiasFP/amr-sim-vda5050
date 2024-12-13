package models

import "time"

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

// Nodes are any point or region within a map
// that represents a place, eg. a charger, a window, a door, a hole.
// Anything really.
func (state *State) FlyTowardsNode(x float64, y float64) {
	if state.AgvPosition.X > x {
		state.AgvPosition.X--
	} else if state.AgvPosition.X < x {
		state.AgvPosition.X++
	}

	if state.AgvPosition.Y > y {
		state.AgvPosition.Y--
	} else if state.AgvPosition.Y < y {
		state.AgvPosition.Y++
	}

}

func GetDefaultState(SN string) State {
	nodeStates := []NodeState{}

	defaultMap := Map{
		MapID:          "99187cd1-8b4b-4f5a-ac11-e455928409de",
		MapVersion:     "0.1.1",
		MapStatus:      "beta",
		MapDescription: "Just a random map",
	}

	state := State{
		HeaderID:           0,
		Timestamp:          time.Now().Format("YYYY-MM-DDTHH:mm:ss.ffZ"),
		Version:            "1.2.3",
		Manufacturer:       "Banana Republic",
		SerialNumber:       SN,
		OrderID:            "",
		OrderUpdateID:      0,
		LastNodeID:         "",
		LastNodeSequenceID: 0,
		NodeStates:         nodeStates,
		EdgeStates:         []EdgeState{},
		Driving:            false,
		ActionStates:       []ActionState{},
		BatteryState: BatteryState{
			BatteryCharge:  99,
			Charging:       false,
			BatteryVoltage: 14,
			BatteryHealth:  100,
			Reach:          0,
		},
		OperatingMode:         "",
		Errors:                []Error{},
		SafetyState:           SafetyState{},
		Maps:                  []Map{defaultMap},
		ZoneSetID:             "",
		Paused:                false,
		NewBaseRequest:        false,
		DistanceSinceLastNode: 0,
		AgvPosition: AgvPosition{
			X:                   150,
			Y:                   150,
			Theta:               0,
			MapID:               defaultMap.MapID,
			PositionInitialized: false,
			MapDescription:      defaultMap.MapDescription,
			LocalizationScore:   0,
			DeviationRange:      0,
		},
		Velocity: Velocity{
			Vx:    10,
			Vy:    10,
			Omega: 0,
		},
		Loads:       []Load{},
		Information: []Info{},
	}
	return state
}
