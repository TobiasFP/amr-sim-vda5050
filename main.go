package main

import (
	"TobiasFP/amrsimvda5050/config"
	"TobiasFP/amrsimvda5050/models"
	"encoding/json"
	"flag"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func main() {
	environment := flag.String("e", "production", "")
	flag.Parse()

	config.Init(*environment)
	conf := config.GetConfig()
	broker := conf.GetString("mqttBroker")
	UniqueSerialNumber := conf.GetBool("uniqueSerialNumber")
	SN := ""
	if UniqueSerialNumber {
		SN = uuid.New().String()
	} else {
		SN = conf.GetString("serialnumber")
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connecting to MQTT broker:", token.Error())
	}

	state := models.GetDefaultState(SN)
	ticksPrMqtt := 1
	tick := 0

	goBack := false
	for {
		// Update time of message
		state.Timestamp = time.Now().Format(time.RFC3339)

		state.BatteryState.BatteryCharge--
		if state.BatteryState.BatteryCharge < 30 {
			state.BatteryState.BatteryCharge = 99
		}
		if !goBack {
			goBack = state.FlyTowardsNode(1000, 350, 5)
		} else {
			goBack = !state.FlyTowardsNode(700, 400, 5)
		}

		time.Sleep(time.Duration(1) * time.Second)
		if tick == 0 {
			message, err := json.Marshal(state)
			if err != nil {
				log.Fatal(err)
			}
			token := client.Publish("state", 0, false, message)
			token.Wait()
		}

		// Looping for the tick.
		tick++
		if tick >= ticksPrMqtt {
			tick = 0
		}
	}
}
