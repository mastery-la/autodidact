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
	transport := homekit.New("12344321", "12345", "./db")

	ch := make(chan homekit.Accessorizable)
	go getAccessories(ch)

	count := 0
	for accessory := range ch {
		count = count + 1
		transport.AddAccessory(accessory)
		if count == 1 {
			close(ch)
		}
	}

	transport.Start()
}

func getAccessories(ch chan homekit.Accessorizable) {
	go func() {
		ts, err := getHoneywellThermostat()
		if err != nil {
			log.Println(err)
		} else {
			ch <- homekit.NewThermostat(ts.Thermostat)
		}
	}()
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
