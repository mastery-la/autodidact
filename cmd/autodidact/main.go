package main

import (
	"log"
	"os"
	"strconv"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"

	"github.com/mastery-la/autodidact/pkg/bridge"

	"github.com/joho/godotenv"
)

func main() {
	info := accessory.Info{
		Name:         "Thermostat",
		SerialNumber: "TCAS3AS14AS3",
		Manufacturer: "Honeywell",
		Model:        "TCC01",
	}
	acc := accessory.NewThermostat(info, 22.0, 15.0, 32.0, 1.0)

	config := hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./.db"}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Panic(err)
	}

	bridge, err := getBridge()
	if err != nil {
		log.Panic(err)
	}

	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) {
		bridge.Thermostat.Thermostat.TargetTemperature.SetValue(temp)
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}

func getBridge() (*bridge.HoneywellThermostat, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DEVICEID, err := strconv.Atoi(os.Getenv("THERMOSTAT_ID"))
	if err != nil {
		log.Panic(err)
	}

	LOGIN := os.Getenv("THERMOSTAT_LOGIN")
	PASSWORD := os.Getenv("THERMOSTAT_PASSWORD")

	return bridge.NewHoneywellThermostat(DEVICEID, LOGIN, PASSWORD)
}
