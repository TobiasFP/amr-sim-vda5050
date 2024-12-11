package main

import (
	"TobiasFP/amrsimvda5050/config"
	"TobiasFP/amrsimvda5050/models"
	"encoding/json"
	"flag"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/exp/rand"
)

func main() {
	environment := flag.String("e", "production", "")
	flag.Parse()

	config.Init(*environment)
	conf := config.GetConfig()
	broker := conf.GetString("mqttBroker")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connecting to MQTT broker:", token.Error())
	}
	nodeStates := []models.NodeState{}

	defaultMap := models.Map{
		MapID:          "99187cd1-8b4b-4f5a-ac11-e455928409de",
		MapVersion:     "0.1.1",
		MapStatus:      "beta",
		MapDescription: "Just a random map",
	}

	state := models.State{
		HeaderID:           0,
		Timestamp:          "2024-11-18T11:40:03.12Z",
		Version:            "1.2.3",
		Manufacturer:       "Banana Republic",
		SerialNumber:       "d9c717cb-158d-4a63-ac2b-45aa8973a5f3",
		OrderID:            "",
		OrderUpdateID:      0,
		LastNodeID:         "",
		LastNodeSequenceID: 0,
		NodeStates:         nodeStates,
		EdgeStates:         []models.EdgeState{},
		Driving:            false,
		ActionStates:       []models.ActionState{},
		BatteryState: models.BatteryState{
			BatteryCharge:  99,
			Charging:       false,
			BatteryVoltage: 14,
			BatteryHealth:  100,
			Reach:          0,
		},
		OperatingMode:         "",
		Errors:                []models.Error{},
		SafetyState:           models.SafetyState{},
		Maps:                  []models.Map{defaultMap},
		ZoneSetID:             "",
		Paused:                false,
		NewBaseRequest:        false,
		DistanceSinceLastNode: 0,
		AgvPosition: models.AgvPosition{
			X:                   150,
			Y:                   150,
			Theta:               0,
			MapID:               defaultMap.MapID,
			PositionInitialized: false,
			MapDescription:      defaultMap.MapDescription,
			LocalizationScore:   0,
			DeviationRange:      0,
		},
		Velocity: models.Velocity{
			Vx:    10,
			Vy:    10,
			Omega: 0,
		},
		Loads:       []models.Load{},
		Information: []models.Info{},
	}

	for {
		xDirection := (rand.Float64() - 0.5) * 20
		yDirection := (rand.Float64() - 0.5) * 20
		state.BatteryState.BatteryCharge--
		if state.BatteryState.BatteryCharge < 30 {
			state.BatteryState.BatteryCharge = 99
		}

		state.AgvPosition.X = state.AgvPosition.X - xDirection
		state.AgvPosition.Y = state.AgvPosition.Y - yDirection

		message, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		token := client.Publish("state", 0, false, message)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}
