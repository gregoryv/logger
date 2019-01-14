[![Build Status](https://travis-ci.org/gregoryv/logger.svg?branch=master)](https://travis-ci.org/gregoryv/logger)
[![codecov](https://codecov.io/gh/gregoryv/logger/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/logger)

[logger](https://godoc.org/github.com/gregoryv/logger) - package provides log.Logger adapters

## Quick start

    go get github.com/gregoryv/logger

In eg. your test

    log := logger.NewTadapter(t)
	// use the logger as you would with log.Logger
	obj := NewSomething(log)
	obj.DoSomething()
	t.Fail()
	// whatever DoSomething() logged ends up using t.Log and t.Logf funcs
