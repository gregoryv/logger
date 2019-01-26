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
	"fmt"
	"log"
	"os"
)

// Returns a logger with log.LstdFlags|log.Lshortfile
func New() Logger {
	return Wrap(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
}

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func Wrap(l *log.Logger) *Wrapped {
	return &Wrapped{l}
}

type Wrapped struct {
	l *log.Logger
}

func (wrap *Wrapped) Log(args ...interface{}) {
	wrap.l.Output(2, fmt.Sprint(args...))
}

func (wrap *Wrapped) Logf(format string, args ...interface{}) {
	wrap.l.Output(2, fmt.Sprintf(format, args...))
}

// Logger that outputs nothing
const Silent silent = iota

type silent int

func (silent) Log(args ...interface{})                 {}
func (silent) Logf(format string, args ...interface{}) {}
