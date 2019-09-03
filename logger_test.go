package logger

import (
	"bytes"
	"log"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_constructor(t *testing.T) {
	l := New()
	assert := asserter.New(t)
	assert(l != nil).Fail()
}

func Test_output_of_logger(t *testing.T) {
	assert := asserter.New(t)
	w := bytes.NewBufferString("")
	l := Wrap(log.New(w, "", log.LstdFlags))

	exp := "emotions assign value to things"
	l.Log(exp)
	assert().Contains(w.Bytes(), exp)

	w.Reset()
	l.Logf("%s", exp)
	assert().Contains(w.Bytes(), exp)

	p := Prefix(l, "p: ")
	w.Reset()
	p.Log("hello")
	assert().Contains(w.Bytes(), "p: hello")

	w.Reset()
	p.Logf("%s", "hello")
	assert().Contains(w.Bytes(), "p: hello")
}

func Test_using_t_as_logger(t *testing.T) {
	type Car struct {
		Logger
	}
	car := &Car{t}
	car.Log("starting")
}
