package controller

import (
	"context"
	"log"
	"math"

	"github.com/mastery-la/autodidact/pkg/util"

	"github.com/go-home-iot/honeywell"
)

type TotalComfortControl struct {
	*Controller
	ts  honeywell.Thermostat
	ctx context.Context
}

func NewTotalComfortControl(ctx context.Context, deviceID int, login string, password string) (*TotalComfortControl, error) {
	tcc := new(TotalComfortControl)

	tcc.Controller = New()
	tcc.ctx = ctx

	tcc.ts = honeywell.NewThermostat(deviceID)
	err := tcc.ts.Connect(tcc.ctx, login, password)
	if err != nil {
		return nil, err
	}

	return tcc, nil
}

func (c *TotalComfortControl) SetCoolF(temp float64) {
	setPoint := math.Round(temp)
	log.Printf("setting cool point to %gºF\n", setPoint)
	c.ts.CoolMode(c.ctx, float32(setPoint), 0)
}

func (c *TotalComfortControl) SetHeatF(temp float64) {
	setPoint := math.Round(temp)
	log.Printf("setting heat point to %gºF\n", setPoint)
	c.ts.HeatMode(c.ctx, float32(setPoint), 0)
}

func (c *TotalComfortControl) SetCoolC(temp float64) {
	log.Printf("setting cool point to %gºC\n", temp)

	fahr := util.Celsius2Fahrenheit(temp)
	c.SetCoolF(fahr)
}

func (c *TotalComfortControl) SetHeatC(temp float64) {
	log.Printf("setting heat point to %gºC\n", temp)

	fahr := util.Celsius2Fahrenheit(temp)
	c.SetHeatF(fahr)
}
