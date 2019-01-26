/*
Package logger adapts gos built in logger.Logger with Log and Logf funcs.

This enables you to write code like

  type Car struct {
     logger.Logger
  }

  log := log.New(os.Stderr, "car: ", log.Lshortfile)
  car := &Car{logger.Adapt(log))

  car.Log("something")
  car.Logf("more %s", here)
*/
package logger

import (
	"io"
	"log"
)

// Returns a logger with log.LstdFlags
func New(w io.Writer) Logger {
	return Adapt(log.New(w, "", log.LstdFlags))
}

type Any interface{}

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

// Adapter maps Log to Print funcs
type Adapter struct{ p Printer }

func (a *Adapter) Log(args ...interface{})            { a.p.Print(args...) }
func (a *Adapter) Logf(f string, args ...interface{}) { a.p.Printf(f, args...) }

func Adapt(p Printer) Logger { return &Adapter{p} }

type Printer interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// Returns a logger with log.LstdFlags|log.Lshortfile
func NewDebug(w io.Writer) Logger {
	return Adapt(log.New(w, "", log.LstdFlags|log.Lshortfile))
}

// Logger that outputs nothing
const Silent silent = iota

type silent int

func (silent) Log(args ...interface{})                 {}
func (silent) Logf(format string, args ...interface{}) {}
