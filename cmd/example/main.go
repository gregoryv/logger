package main

import (
	"github.com/gregoryv/logger"
)

func main() {
	type Car struct {
		logger.Logger
	}

	car := &Car{logger.New()}
	car.Log("brakes are failing")
	car.Logf("reached speed limit %v", 100)
}
