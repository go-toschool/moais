package main

import (
	"fmt"

	"github.com/WiseGrowth/go-wisebot/rasp"
)

func getNetworks() ([]*rasp.Network, error) {
	networks, err := rasp.AvailableNetworks()
	if err != nil {
		return nil, err
	}

	return networks, nil
}

func setNetwork(ESSID string, password string) error {
	// define network
	network := new(rasp.Network)
	network.ESSID = ESSID
	network.Password = password

	fmt.Sprintf("ESSID: %q Password: %q", network.ESSID, network.Password)

	// set connecting network color
	err := led.setYellowColor()
	if err != nil {
		return err
	}

	// set new wifi params
	err = rasp.SetupWifi(network)
	if err != nil {
		// set error network color
		e := led.setRedColor()
		if e != nil {
			return err
		}

		if err == rasp.ErrNoWifi {
			fmt.Println("No Wifi", err.Error())
			return err
		}

		return err
	}

	// set connected network color
	if err := led.setGreenColor(); err != nil {
		return err
	}

	fmt.Println("Wifi Connected")
	return nil
}
