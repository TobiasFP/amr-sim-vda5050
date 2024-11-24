package main

import (
	"TobiasFP/amrsimvda5050/config"
	"TobiasFP/amrsimvda5050/models"
	"encoding/json"
	"flag"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	state := models.State{
		HeaderID:              0,
		Timestamp:             "2024-11-18T11:40:03.12Z",
		Version:               "1.2.3",
		Manufacturer:          "BananaRepublic",
		SerialNumber:          "d9c717ca-158d-4a63-ac2b-45aa8973a5f3",
		OrderID:               "",
		OrderUpdateID:         0,
		LastNodeID:            "",
		LastNodeSequenceID:    0,
		NodeStates:            nodeStates,
		EdgeStates:            []models.EdgeState{},
		Driving:               false,
		ActionStates:          []models.ActionState{},
		BatteryState:          models.BatteryState{},
		OperatingMode:         "",
		Errors:                []models.Error{},
		SafetyState:           models.SafetyState{},
		Maps:                  []models.Map{},
		ZoneSetID:             "",
		Paused:                false,
		NewBaseRequest:        false,
		DistanceSinceLastNode: 0,
		AgvPosition:           models.AgvPosition{},
		Velocity:              models.Velocity{},
		Loads:                 []models.Load{},
		Information:           []models.Info{},
	}
	message, err := json.Marshal(state)
	if err != nil {
		log.Fatal(err)
	}
	for {
		token := client.Publish("state", 0, false, message)
		token.Wait()
		time.Sleep(15 * time.Second)
	}
}
