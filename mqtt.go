package main

import (
	"fmt"
	"time"

	"github.com/WiseGrowth/go-wisebot/config"
	"github.com/WiseGrowth/go-wisebot/rasp"
	"github.com/WiseGrowth/wisebot-operator/iot"
	"github.com/yanzay/log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// MqttClient ...
type MqttClient struct {
	client *iot.Client
}

//NewMQTTClient ...
func NewMQTTClient(config *config.ConfigV2) (*MqttClient, error) {
	// Define certs to connect to broker
	cert, err := config.GetTLSCertificate()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Define MQTT client
	client, err := iot.NewClient(
		iot.SetHost(config.IOTHost),
		iot.SetCertificate(*cert),
		iot.SetClientID("op-"+config.WisebotID),
	)

	return &MqttClient{
		client: client,
	}, nil
}

func (m *MqttClient) subscribeDefaultTopics() error {
	log.Debug("Subscribing topics")
	// if err := m.client.Subscribe("/storage/"+ID+"/healthz", healthzMQTTHandler); err != nil {
	// 	return err
	// }

	return nil
}

// SubscribeTopic ...
func (m *MqttClient) SubscribeTopic(topic string, handler MQTT.MessageHandler) error {
	fmt.Println("Subscribing " + topic)

	if err := m.client.Subscribe(topic, handler); err != nil {
		return err
	}

	return nil
}

// Serve ...
func (m *MqttClient) Serve() error {
	fmt.Println("Connecting to MQTT Broker")

	if err := m.client.Connect(); err != nil {
		return err
	}

	// Set green color when mqtt is connected
	err := led.setGreenColor()
	if err != nil {
		return err
	}

	// TODO(ca): if mqtt is reconnecting set led to yellow color
	// TODO(ca): get ${API_URL}/wisebots/get-info again when mqtt is reconnected
	// TODO(ca): if mqtt is desconnected set led to red color

	// Get network connection status
	isConnected, err := rasp.IsConnected()
	if err != nil {
		return err
	}

	// Check connection and connect mqtt server. Otherwise retry connection.
	if isConnected {
		// Subscribe to wisebot topics
		err := m.subscribeDefaultTopics()
		if err != nil {
			return err
		}

		fmt.Println("Running MQTT Service")
	} else {
		tick := time.NewTicker(15 * time.Second)

		defer tick.Stop()

		for range tick.C {
			//TODO(ca): add limit to retry connections
			fmt.Println("Trying connect to MQTT Service")
			isConnected, _ := rasp.IsConnected()
			if err != nil {
				//TODO(ca):
				return err
			}

			if isConnected {
				err := m.subscribeDefaultTopics()
				if err != nil {
					return err
				}

				fmt.Println("Running MQTT Service")
			}
		}
	}

	return nil
}
