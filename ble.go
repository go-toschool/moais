package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/WiseGrowth/go-wisebot/config"
	"github.com/WiseGrowth/go-wisebot/rasp"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/paypal/gatt/examples/service"
)

const (
	//TODO(ca): read manifest
	serviceUUID = "ffffffff-ffff-ffff-ffff-fffffffffff1"

	getNetworksUUID     = "ffffffff-ffff-ffff-ffff-fffffffffff2"
	setNetworkUUID      = "ffffffff-ffff-ffff-ffff-fffffffffff3"
	closeConnectionUUID = "ffffffff-ffff-ffff-ffff-fffffffffff4"

	packerEndMarker      = "__wise__"
	packerErrorMarker    = "__error__"
	blePacketHeaderBytes = 3
	notifyIntervalTime   = 10
)

// WisebotBle ...
type WisebotBle struct {
	name            string
	device          gatt.Device
	IsAdvertisement bool
	IsConnected     bool
}

// NewWisebotBle ...
func NewWisebotBle(config *config.ConfigV2) (*WisebotBle, error) {
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		fmt.Println("Failed to open device, err: ", err.Error())
		return nil, err
	}

	return &WisebotBle{
		name:            config.BLE.Name,
		device:          d,
		IsAdvertisement: false,
		IsConnected:     false,
	}, nil
}

// BaseService ...
func BaseService() *gatt.Service {
	s := gatt.NewService(gatt.MustParseUUID(serviceUUID))

	s.AddCharacteristic(gatt.MustParseUUID(setNetworkUUID)).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Wrote:", string(data))

			var network rasp.Network
			if err := json.Unmarshal(data, &network); err != nil {
				fmt.Println("setNetworkUUID Unmarshal", err.Error())
				return gatt.StatusUnexpectedError
			}

			err := setNetwork(network.ESSID, network.Password)
			if err != nil {
				fmt.Println("Error while set wifi network", err.Error())
				return gatt.StatusUnexpectedError
			}

			return gatt.StatusSuccess
		},
	)

	s.AddCharacteristic(gatt.MustParseUUID(closeConnectionUUID)).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
			log.Println("Wrote:", string(data))

			if err := led.setRedColor(); err != nil {
				fmt.Println("led.setRedColor", err.Error())
				return gatt.StatusUnexpectedError
			}

			// ENV_FIRST_TIME = TRUE

			// if err := ble.Disconnect(); err != nil {
			// 	fmt.Println("BLE Disconnect", err.Error())
			// 	return gatt.StatusUnexpectedError
			// }

			if err := ble.StopAdvertising(); err != nil {
				fmt.Println("BLE Stop Advertising", err.Error())
				return gatt.StatusUnexpectedError
			}

			// stop core service
			// start core service

			return gatt.StatusSuccess
		},
	)

	s.AddCharacteristic(gatt.MustParseUUID(getNetworksUUID)).HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
			networks, err := getNetworks()
			if err != nil {
				fmt.Println("getNetworks", err.Error())
				n.Write([]byte(packerErrorMarker))
				n.Done()
				return
			}

			networksBytes, err := json.Marshal(networks)
			if err != nil {
				fmt.Println("json.Marshal(networks)", err.Error())
				n.Write([]byte(packerErrorMarker))
				n.Done()
				return
			}

			fullNetworksBytes := []byte(networksBytes)
			mtu := r.Central.MTU()
			bufferLen := mtu - blePacketHeaderBytes
			sliceLen := int(math.Ceil(float64(len(fullNetworksBytes)) / float64(bufferLen)))

			if len(fullNetworksBytes) <= bufferLen {
				n.Write(fullNetworksBytes)
				return
			}

			offset := 0
			for i := 0; i < sliceLen; i++ {
				data := fullNetworksBytes[offset : offset+sliceLen]
				n.Write(data)
				offset += sliceLen

				time.Sleep(notifyIntervalTime * time.Millisecond)
			}

			fmt.Println("wifi network list was sended successfully")
		},
	)

	return s
}

// Start ...
func (wb *WisebotBle) Start() error {
	//register handlers
	wb.device.Handle(
		gatt.CentralConnected(wb.connectedHandler),
		gatt.CentralDisconnected(wb.disconnectedHandler),
	)

	wb.device.Init(wb.stateHandler)

	// TODO(ca): remove this, don't sure if this is neccesary
	select {}

	return nil
}

func (wb *WisebotBle) stateHandler(d gatt.Device, s gatt.State) {
	fmt.Printf("State: %s\n", s)
	switch s {
	case gatt.StatePoweredOn:
		// Set default gatt service
		d.AddService(service.NewGapService("wisebot-" + wisebotID))
		d.AddService(service.NewGattService())

		// Wisebot base BLE service
		s1 := BaseService()
		d.AddService(s1)

		// Advertise device name and service's UUIDs.
		d.AdvertiseNameAndServices("wisebot-"+wisebotID, []gatt.UUID{s1.UUID()})

	default:
	}
}

func (wb *WisebotBle) connectedHandler(c gatt.Central) {
	fmt.Println("Connect: ", c.ID())
	wb.IsConnected = true
}

func (wb *WisebotBle) disconnectedHandler(c gatt.Central) {
	fmt.Println("Disconnect: ", c.ID())
	wb.IsConnected = false
}

// StopAdvertising ...
func (wb *WisebotBle) StopAdvertising() error {
	if err := wb.device.StopAdvertising(); err != nil {
		return err
	}

	return nil
}
