/* Package logger adapts log.Logger with Log and Logf funcs.

There are two benefits of this, the smaller Logger interface limits
logging to output only and can easily be replaced by the testing.T
object during testing.

	type Car struct {
		logger.Logger
	}

	car := &Car{logger.New()}
	car.Log("brakes are failing")
	car.Logf("reached speed limit %s", 100)
*/
package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

// Returns a logger with log.LstdFlags|log.Lshortfile that writes to stderr
func New() *Wrapped {
	return Wrap(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
}

// Returns a logger with empty flags that writes to stdout
func NewProgress() *Wrapped {
	return Wrap(log.New(os.Stdout, "", 0))
}

type Wrapped struct {
	l *log.Logger
}

// Wrap creates and adapter from Logger to log.Logger
func Wrap(l *log.Logger) *Wrapped {
	return &Wrapped{l}
}

func (wrap *Wrapped) Log(args ...interface{}) {
	wrap.l.Output(2, fmt.Sprint(args...))
}

func (wrap *Wrapped) Logf(format string, args ...interface{}) {
	wrap.l.Output(2, fmt.Sprintf(format, args...))
}

// Logger that outputs nothing
var Silent Logger = silent(0)

type silent int

func (silent) Log(args ...interface{})                 {}
func (silent) Logf(format string, args ...interface{}) {}

func Prefix(wrap *Wrapped, prefix string) *Prefixed {
	return &Prefixed{wrap, prefix}
}

type Prefixed struct {
	*Wrapped
	prefix string
}

func (p *Prefixed) Log(args ...interface{}) {
	p.l.Output(2, p.prefix+fmt.Sprint(args...))
}

func (p *Prefixed) Logf(format string, args ...interface{}) {
	p.l.Output(2, fmt.Sprintf(p.prefix+format, args...))
}
