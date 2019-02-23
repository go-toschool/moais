package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/WiseGrowth/go-wisebot/config"
	"github.com/WiseGrowth/go-wisebot/logger"
)

var (
	button *WisebotButton
	led    *WisebotLed
	ble    *WisebotBle

	wisebotConfig *config.ConfigV2
	mqttClient    *MqttClient
)

const (
	coreURL        = "http://localhost:5005/"
	iotHost        = "mqtt://localhost"
	coreHealthzURL = coreURL + "healthz"

	wisebotID         = "asdf1234"
	wisebotConfigPath = "~/.config/wisebot/config.json"
)



// Wisebot ...
type Wisebot struct {
	ID string

	GlobalValues WisebotGlobalValues

	// private rulesEngine: RulesEngine
	// private devicesEngine: DevicesEngine
	// private talkative: Talkative
	// private globalValues: IGlobalValues = {}
	// private getInfoInterval: NodeJS.Timer = null
	// private getInfoSubscribedId: string[] = []
}

//BalenaENV ...
type BalenaENV struct {
	isBluetoothMode bool
}

func main() {
	// Load wisebot config v2
	wisebotConfig, err := config.LoadConfigV2(wisebotConfigPath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Define quit chan to send interruption
	quit := make(chan struct{})

	//TODO(ca): remove this when have BALENA_ENV values
	balenaENV := new(BalenaENV)
	balenaENV.isBluetoothMode = true

	// Initialize http server
	// err := NewHTTPServer()

	// httpServer = httpStorage.NewHTTPServer(httpCtx)
	// log.Info("Running server on: " + httpServer.Addr)
	// go func() {
	// 	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Error(err)
	// 		quit <- struct{}{}
	// 		return
	// 	}
	// }()

	// Initialize wisebot led service
	led = NewWisebotLed()
	go func() {
		fmt.Println("here in led goroutine")
		err := led.Start()
		if err != nil {
			//TODO(ca): send to channel
		}
	}()

	// Initialize wisebot button service
	button = NewWisebotButton()
	go func() {
		fmt.Println("here in button goroutine")
		err := button.Start()
		if err != nil {
			//TODO(ca): send to quit channel
		}
	}()

	// Initialize wisebot ble service
	ble, err := NewWisebotBle(wisebotConfig)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		fmt.Println("here in ble goroutine")
		err := ble.Start()
		if err != nil {
			//TODO(ca): send to quit channel
		}
	}()

	// Initialize MQTT Client
	mqttClient, err := NewMQTTClient(wisebotConfig)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		err := mqttClient.Serve()
		if err != nil {
			//TODO(ca): send to channel
		}
	}()

	listenInterrupt(quit)
	<-quit
	gracefullShutdown()
}

func listenInterrupt(quit chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		fmt.Println("Signal received - " + s.String())
		quit <- struct{}{}
	}()
}

func gracefullShutdown() {
	log := logger.GetLogger()

	mqttClient.client.Disconnect(250)
	log.Info("[MQTT] Disconnected")

	log.Debug("Gracefully shutdown")

	// if err := httpServer.Shutdown(nil); err != nil {
	// 	log.Error(err.Error())
	// }
}
