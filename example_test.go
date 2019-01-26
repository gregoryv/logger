package logger_test

import (
	"github.com/gregoryv/logger"
)

func ExampleNew() {
	type Car struct {
		logger.Logger
	}

	car := &Car{logger.New("car: ")}
	car.Log("brakes are failing")
	car.Logf("reached speed limit %s", 100)
}
