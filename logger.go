/* Package logger adapts log.Logger with Log and Logf funcs.

There are two benefits of this, the smaller Logger interface limits
logging to output only and can easily be replaced by the testing.T
object during testing.

	type Car struct {
		logger.Logger
	}

	car := &Car{logger.New("car: ")}
	car.Log("brakes are failing")
	car.Logf("reached speed limit %s", 100)
*/
package logger

import (
	"log"
	"os"
)

// Returns a logger with log.LstdFlags|log.Lshortfile
func New(prefix string) Logger {
	return Adapt(log.New(os.Stderr, prefix, log.LstdFlags|log.Lshortfile))
}

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

// Logger that outputs nothing
const Silent silent = iota

type silent int

func (silent) Log(args ...interface{})                 {}
func (silent) Logf(format string, args ...interface{}) {}
