package logger_test

import (
	"github.com/gregoryv/logger"
)

func ExampleNew() {
	type Car struct {
		logger.Logger
	}

	car := &Car{logger.New()}
	car.Log("brakes are failing")
	car.Logf("reached speed limit %v", 100)
	// output:
}

func Example_silent() {
	l := logger.Silent
	l.Log("nada")
	l.Logf("%s", "nada")
	// output:
}
