package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-home-iot/honeywell"
	"github.com/joho/godotenv"
)

func main() {
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

	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)
	err = ts.Connect(ctx, LOGIN, PASSWORD)
	if err != nil {
		log.Panic(err)
	}

	status, err := ts.FetchStatus(ctx)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Currently %gºF", status.LatestData.UIData.DispTemperature)

}
