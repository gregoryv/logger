[![Build Status](https://travis-ci.org/gregoryv/logger.svg?branch=master)](https://travis-ci.org/gregoryv/logger)
[![codecov](https://codecov.io/gh/gregoryv/logger/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/logger)

[logger](https://godoc.org/github.com/gregoryv/logger) - package provides log.Logger adapters

## Quick start

    go get github.com/gregoryv/logger

Use in your code

    type Car struct {
        logger.Logger
    }

    car := &Car{logger.New("car: ")}
    car.Log("brakes are failing")
    car.Logf("reached speed limit %s", 100)

or when testing the logger can be replaced by testing.T

    func Test_car_speed(t *testing.T) {
        car := &Car{t}
        // asserts here
    }
