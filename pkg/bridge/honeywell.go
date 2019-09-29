package bridge

import (
	"context"
	"log"
	"strconv"

	"github.com/mastery-la/autodidact/pkg/attribute"
	"github.com/mastery-la/autodidact/pkg/controller"

	"github.com/mastery-la/autodidact/pkg/node"
)

type HoneywellThermostat struct {
	*Bridge
	*node.Thermostat
	controller *controller.TotalComfortControl
}

func NewHoneywellThermostat(deviceID int, login string, password string) (*HoneywellThermostat, error) {
	hw := new(HoneywellThermostat)

	ts := node.NewThermostat(strconv.Itoa(deviceID))
	hw.Bridge = New(ts.Node)

	hw.Thermostat = ts

	ctx := context.TODO()
	tcc, err := controller.NewTotalComfortControl(ctx, deviceID, login, password)
	if err != nil {
		return nil, err
	}

	hw.controller = tcc
	hw.AddController(tcc.Controller)

	hw.setupBridge()

	return hw, nil
}

func (b *HoneywellThermostat) setupBridge() {
	b.Thermostat.Thermostat.
		TargetTemperature.
		OnChange(func(a *attribute.Attribute, newValue, oldValue interface{}) {
			if newValue == oldValue {
				return
			}

			temp := newValue.(float64)
			log.Printf("client wants %g", temp)
			b.controller.SetCoolC(temp)
		})
}
