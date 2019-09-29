package main

import (
	"log"
	"os"
	"strconv"

	"github.com/mastery-la/autodidact/pkg/protocol/homekit"

	"github.com/mastery-la/autodidact/pkg/bridge"

	"github.com/joho/godotenv"
)

func main() {
	honeywell, err := getHoneywellThermostat()
	if err != nil {
		log.Panic(err)
	}

	transport := homekit.New("12344321", "12345", "./db")

	transport.AddAccessory(homekit.NewThermostat(honeywell.Thermostat))

	transport.Start()
}

func getHoneywellThermostat() (*bridge.HoneywellThermostat, error) {
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
